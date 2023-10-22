package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kevinronu/fibonacci-tg-bot/utils"
)

func main() {
	portString := utils.GetEnv("SERVER_PORT")
	// Create a custom ServeMux to register our handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleTelegramWebHook)

	// Create an HTTP server
	srv := &http.Server{
		Addr:    ":" + portString,
		Handler: mux,
	}

	// Start the HTTP server
	log.Println("INFO: Server is starting on port:", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("FATAL: Failed to start server:", err)
	}
}

func handleTelegramWebHook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleTelegramWebHook")
}
