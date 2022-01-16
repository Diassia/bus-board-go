package main

import (
	"bus-stops/stops"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type stopHandler struct {
	format string
}

type postcodeHandler struct {
	postcode string
}

func (sh stopHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sm := sh.format
	w.Write([]byte("The closest bus stop is: " + sm))
}

func (ph postcodeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pc := r.URL.Query().Get("postcode")
	long, lat := stops.FetchLocation(pc)

	w.Write([]byte("Postcode " + strings.ToUpper(pc) + " is at longitude: " + fmt.Sprintf("%f", long) + ". Latitude: " + fmt.Sprintf("%f", lat) + "\n"))
	w.Write([]byte(stops.FetchNearbyStops(long, lat)))
}

func main() {
	// http.Handle("/stop", stopHandler{format: "Ruskington B17"})
	http.Handle("/stop", postcodeHandler{})
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
