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
  "suggestedFilename": "W409101448A_2025_06_18_08_51_12.zip",
  "files": [
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED9A8AE4899E53AAC0C/TI01442TPT_0118.pdf",
      "zipPath": "Technical_information/PT/TI01442TPT_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED9A8AE49FA1B53EC0E/TI01442TES_0118.pdf",
      "zipPath": "Technical_information/ES/TI01442TES_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED9B2AFD71708DE363F/TI01442TRU_0118.pdf",
      "zipPath": "Technical_information/RU/TI01442TRU_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED9A8AE4964BABF2C0E/TI01442TFR_0118.pdf",
      "zipPath": "Technical_information/FR/TI01442TFR_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED9A8AE4AB73422CC11/TI01442TJA_0118.pdf",
      "zipPath": "Technical_information/JA/TI01442TJA_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED997E92C8D8CDCD5BE/TI01442TDE_0118.pdf",
      "zipPath": "Technical_information/DE/TI01442TDE_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDAA4EA81BC1A6F0ED9/TI01442TPL_0118.pdf",
      "zipPath": "Technical_information/PL/TI01442TPL_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED997E92DB88877F5C1/TI01442TEN_0118.pdf",
      "zipPath": "Technical_information/EN/TI01442TEN_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED9B2AFD7E6E4125642/TI01442TIT_0118.pdf",
      "zipPath": "Technical_information/IT/TI01442TIT_0118.pdf"
    }
  ]
}`), os.Getenv("ZIP_STREAMER_KEY_PHRASE")))
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
