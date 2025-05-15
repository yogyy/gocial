package main

import (
	"log"

	"github.com/gocial/internal/env"
)

func main() {
	app := &application{
		config{addr: env.GetString("ADDR", ":8080")},
	}
	mux := app.mount()

	log.Fatal(app.run(mux))
}
