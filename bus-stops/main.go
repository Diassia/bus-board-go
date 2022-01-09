package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type PostcodeData struct {
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}

type Postcoderesponse struct {
	Result PostcodeData
}

// type Stopsresponse struct {
// 	Result
// }

func FetchLocation(postcode string) (float32, float32) {
	URL := "https://api.postcodes.io/postcodes/" + postcode

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var pResp Postcoderesponse
	error := json.Unmarshal(bodyBytes, &pResp)
	if error != nil {
		log.Println(err)
	}
	longitude := pResp.Result.Longitude
	latitude := pResp.Result.Latitude

	return longitude, latitude
}

func FetchNearbyStops(long float32, lat float32) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// expire 30 days (from 09/01)
	id := os.Getenv("TRANSPORT_ID")
	key := os.Getenv("API_KEY")

	URL := "http://transportapi.com/v3/uk/places.json?lat=" + fmt.Sprintf("%f", lat) + "&lon=" + fmt.Sprintf("%f", long) + "&type=bus_stop&app_id=" + id + "&app_key=" + key

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	s := string(bodyBytes)
	fmt.Println(s)

}

func main() {
	longitude, latitude := FetchLocation("nw51tl")
	FetchNearbyStops(longitude, latitude)

	fmt.Printf("Longitude: %v\nLatitude: %v\n", longitude, latitude)
}
