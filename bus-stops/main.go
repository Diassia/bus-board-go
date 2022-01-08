package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type PostcodeData struct {
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}

type Postcoderesponse struct {
	Result PostcodeData
}

func FetchPostcodeCoords(postcode string) (float32, float32) {
	URL := "https://api.postcodes.io/postcodes/" + postcode
	fmt.Println(URL)

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

func main() {
	longitude, latitude := FetchPostcodeCoords("nw51tl")

	fmt.Printf("Longitude: %v\nLatitude: %v\n", longitude, latitude)
}
