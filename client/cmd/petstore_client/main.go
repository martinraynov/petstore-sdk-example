package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"petstore_client/router"
	"runtime"

	"github.com/rs/cors"
)

// setupGlobalMiddleware will setup CORS
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler
	return handleCORS(handler)
}

// endpointListenAndServe is the http/https listener
func endpointListenAndServe() {
	router := router.NewRouter(true)

	go func() {
		log.Fatal(http.ListenAndServe(":81", setupGlobalMiddleware(router)))
	}()
}

func helpRequest() {
	fmt.Println("[HELP] GO Runtime Version : ", runtime.Version())
	os.Exit(1)
}

// Main function for the Hofundapi application
func main() {
	// If argument is passed only display info and shutdown
	if os.Args != nil && len(os.Args) > 1 {
		switch os.Args[1] {
		case "help":
			helpRequest()
			break
		default:
			fmt.Println("Unknown argument :", os.Args[1])
			os.Exit(1)
		}
	}

	// Init api listener
	endpointListenAndServe()

	mainChannel := make(chan bool)
	<-mainChannel
}
