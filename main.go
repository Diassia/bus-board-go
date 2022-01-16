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

	mux := http.NewServeMux()
	sh := stopHandler{format: "Ruskington B17"}
	mux.Handle("/stop", sh) //this will handle the request and call our handle function
	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}