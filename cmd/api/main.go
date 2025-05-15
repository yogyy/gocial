package main

import "log"

func main() {
	app := &application{
		config{addr: ":8080"},
	}
	mux := app.mount()

	log.Fatal(app.run(mux))
}
