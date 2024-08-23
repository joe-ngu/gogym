package main

import (
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
	server.Run()
}
