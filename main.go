package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	zip_streamer "github.com/scosman/zipstreamer/zip_streamer"
)

func main() {
	fmt.Println(zip_streamer.EncryptIt([]byte(`{
  "suggestedFilename": "W409101448A_2025_06_18_09_04_27.zip",
  "files": [
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EED8FE7EA7AAFEDD16F/TI01392TRU_0522-00.pdf",
      "zipPath": "Technical_information/RU/TI01392TRU_0522-00.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDA9B9B5F06DBE88985/BA01915TEL_0118.pdf",
      "zipPath": "Operating_instruction/EL/BA01915TEL_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EEDBBB0339BB08E9292/BA01915TZH_0118-00.pdf",
      "zipPath": "Operating_instruction/ZH/BA01915TZH_0118-00.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED997E92C8D8CDCD5BE/TI01442TDE_0118.pdf",
      "zipPath": "Technical_information/DE/TI01442TDE_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED9B88723E5485F2BAC/KA01414TJA_0219.pdf",
      "zipPath": "Short_operating_instruction/JA/KA01414TJA_0219.pdf"
    }]}`), os.Getenv("ZIP_STREAMER_KEY_PHRASE")))
	zipServer := zip_streamer.NewServer()
	zipServer.Compression = (os.Getenv("ZS_COMPRESSION") == "DEFLATE")
	zipServer.ListfileUrlPrefix = os.Getenv("ZS_LISTFILE_URL_PREFIX")

	port := os.Getenv("PORT")
	if port == "" {
		port = "4008"
	}

	httpServer := &http.Server{
		Addr:        ":" + port,
		Handler:     zipServer,
		ReadTimeout: 10 * time.Second,
	}

	shutdownChannel := make(chan os.Signal, 10)
	go func() {
		log.Printf("Server starting on port %s", port)
		err := httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Printf("Server Error: %s", err)
		}
		shutdownChannel <- syscall.SIGUSR1
	}()

	// Listen for os signal for graceful shutdown
	signal.Notify(shutdownChannel, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	// Wait for shutdown signal, then shut down
	shutdownSignal := <-shutdownChannel
	log.Printf("Received signal (%s), shutting down...", shutdownSignal.String())
	httpServer.Shutdown(context.Background())

	// Exit was not expected, return non 0 exit code
	if shutdownSignal == syscall.SIGUSR1 {
		os.Exit(1)
	}
}
