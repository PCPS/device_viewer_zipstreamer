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
  "suggestedFilename": "M3103D05G00_2025_06_18_08_58_11.zip",
  "files": [
    {
      "url": "https://bdih-download.endress.com/files/DRAWXL/0200030000091EE780CB28DE020FA596/CM42-OAA100EAD01!71329995!D!160603.PDF",
      "zipPath": "Operating_instruction/DE/CM42-OAA100EAD01_71329995_D_160603.PDF"
    },
    {
      "url": "https://bdih-download.endress.com/files/DRAWXL/0200030000091ED785BD6A5F937C06C3/M3103D05G00_PFCC_A2_71212267.pdf",
      "zipPath": "Factory_calibration_certificate/EN/M3103D05G00_PFCC_A2_71212267.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DRAWXL/0200030000091ED785BD6A5F937C06C3/M3103D05G00_PFCC_A2_71212267.pdf",
      "zipPath": "Factory_calibration_certificate/DE/M3103D05G00_PFCC_A2_71212267.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DRAWXL/0200030000091EE780CB290467CDC596/CM42-OAA100EAD01!71329994!D!140616.PDF",
      "zipPath": "Operating_instruction/DE/CM42-OAA100EAD01_71329994_D_140616.PDF"
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
