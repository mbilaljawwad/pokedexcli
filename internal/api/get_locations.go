package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/mbilaljawwad/pokedexcli/internal/modal"
)

func GetLocations (url string) ( modal.Locations) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}	
	defer res.Body.Close()
	
	body, err := io.ReadAll(res.Body)
	if (err != nil) {
		log.Fatal(err)
	}
	// it would be nice to log proper status codes errors here
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	locations := modal.Locations{}

	err = json.Unmarshal(body, &locations)
	if err !=	nil {
		log.Fatal(err)
	}

	return locations
}