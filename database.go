package main

import (
	"log"

	"github.com/sdomino/scribble"
)

func initDB() *scribble.Driver {
	db, err := scribble.New("./data", nil)
	if err != nil {
		log.Fatal("Error initialising the database", err)
	}
	return db
}
