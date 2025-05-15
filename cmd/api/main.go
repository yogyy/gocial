package main

import "log"

func main() {
	app := &application{
		config{addr: ":3000"},
	}
	mux := app.mount()

	log.Fatal(app.run(mux))
}
