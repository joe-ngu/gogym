package main

import (
	"log"

	"github.com/joe-ngu/gogym/store"
	"github.com/joe-ngu/gogym/utils"
)

func main() {
	utils.LoadDotenv()

	store, err := store.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created new PostgresStore")

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":3000", store)
	server.Run()
}
