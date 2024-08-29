package main

import (
	"fmt"
	"log"

	"github.com/joe-ngu/gogym/store"
)

func main() {
	store, err := store.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created new PostgresStore")
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":8000", store)
	server.setupRoutes()

	fmt.Println("Server listening on port :8000")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
