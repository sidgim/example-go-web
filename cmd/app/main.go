package main

import (
	"log"
)

func main() {
	srv, err := InitializeApp()
	if err != nil {
		log.Fatalf("failed to initialize: %v", err)
	}
	log.Println("🚀 Server listening on", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server stopped with error: %v", err)
	}
}
