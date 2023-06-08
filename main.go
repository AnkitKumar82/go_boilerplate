package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/AnkitKumar82/go_boilerplate/controllers"
)

func main() {
	fmt.Println("Initializing server...")

	port := os.Getenv("PORT")
	log.Printf("server started on %v", port)

	if port == "" {
		port = ":8080"
	}

	mux := http.NewServeMux()

	controllers.InitialzeControllers(mux)

	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	log.Printf("server started on %v", port)
	log.Fatal(server.ListenAndServe())
}
