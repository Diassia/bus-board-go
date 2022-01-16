package stops

import (
	"bytes"
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

type BusStop struct {
	Name      string  `json:"name"`
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
	Distance  float32 `json:"distance"`
	Atcocode  string  `json:"atcocode"`
}

type Postcoderesponse struct {
	Result PostcodeData
}

type Stopsresponse struct {
	Member []BusStop
}

func PrettyJSON(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

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

	var sResp Stopsresponse
	error := json.Unmarshal(bodyBytes, &sResp)
	if error != nil {
		log.Println(err)
	}

	// bus 1
	b1name := sResp.Member[0].Name
	b1long := sResp.Member[0].Longitude
	b1lat := sResp.Member[0].Latitude
	fmt.Printf("%v is at longitude: %v and latitude: %v\n", b1name, b1long, b1lat)

	// bus 2
	b2name := sResp.Member[1].Name
	b2long := sResp.Member[1].Longitude
	b2lat := sResp.Member[1].Latitude
	fmt.Printf("%v is at longitude: %v and latitude: %v\n", b2name, b2long, b2lat)

	res, err := PrettyJSON(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

func main() {
	fmt.Println("Enter a postcode: ")
	var postcode string
	fmt.Scanln(&postcode)

	longitude, latitude := FetchLocation(postcode)
	FetchNearbyStops(longitude, latitude)
}
