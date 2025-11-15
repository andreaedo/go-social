package main

import (
	"log"

	"github.com/andreaedo/go-social/internal/env"
	"github.com/andreaedo/go-social/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	storage := store.NewStorage(nil)

	app := &application{
		config:  cfg,
		storage: storage,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))

}
