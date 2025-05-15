package main

import "net/http"

func (app *application) healthChekHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
