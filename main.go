package main

import (
	"log"
	"net/http"
)

type stopHandler struct {
	format string
}

func (sh stopHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sm := sh.format
	w.Write([]byte("The closest bus stop is: " + sm))
}

func main() {
	http.Handle("/stop", stopHandler{format: "Ruskington B17"})
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}