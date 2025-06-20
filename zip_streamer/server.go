package zip_streamer

import (
	"archive/zip"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	router            *mux.Router
	linkCache         LinkCache
	Compression       bool
	ListfileUrlPrefix string
}

func NewServer() *Server {
	r := mux.NewRouter()

	timeout := time.Second * 60
	server := Server{
		router:      r,
		linkCache:   NewLinkCache(&timeout),
		Compression: false,
	}

	r.HandleFunc("/download", server.HandlePostDownload).Methods("POST")

	return &server
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	originsOk := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	handlers.CORS(originsOk, headersOk, methodsOk)(s.router).ServeHTTP(w, r)
}

func (s *Server) parseZipRequest(w http.ResponseWriter, req *http.Request) (*ZipDescriptor, error) {
	keyPhrase := os.Getenv("ZIP_STREAMER_KEY_PHRASE")
	body, err := ioutil.ReadAll(req.Body)
	decryptedBody := DecryptIt(string(body), keyPhrase)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status":"error","error":"missing body"}`))
		return nil, err
	}

	ZipDescriptor, err := UnmarshalJsonZipDescriptor([]byte(decryptedBody))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status":"error","error":"invalid body"}`))
		return nil, err
	}

	return ZipDescriptor, nil
}

func (s *Server) HandlePostDownload(w http.ResponseWriter, req *http.Request) {
	zipDescriptor, err := s.parseZipRequest(w, req)
	if err != nil {
		return
	}

	s.streamEntries(zipDescriptor, w)
}

func (s *Server) streamEntries(zipDescriptor *ZipDescriptor, w http.ResponseWriter) {
	zipStreamer, err := NewZipStream(zipDescriptor.Files(), w)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status":"error","error":"invalid entries"}`))
		return
	}

	if s.Compression {
		zipStreamer.CompressionMethod = zip.Deflate
	}

	// need to write the header before bytes
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", zipDescriptor.EscapedSuggestedFilename()))
	w.WriteHeader(http.StatusOK)
	err = zipStreamer.StreamAllFiles()

	if err != nil {
		// Close the connection so the client gets an error instead of 200 but invalid file
		closeForError(w)
	}
}

func closeForError(w http.ResponseWriter) {
	hj, ok := w.(http.Hijacker)

	if !ok {
		return
	}

	conn, _, err := hj.Hijack()
	if err != nil {
		return
	}

	conn.Close()
}
