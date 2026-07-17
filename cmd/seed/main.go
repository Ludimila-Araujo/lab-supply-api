package main

import (
	"log"

	"github.com/Ludimila-Araujo/lab-supply-api/internal/config"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/database"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/seed"
)

func main() {

	cfg := config.Load()

	db, err := database.NewConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := seed.Run(db); err != nil {
		log.Fatal(err)
	}
}
