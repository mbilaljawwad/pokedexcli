package cmd

import (
	"fmt"

	"github.com/mbilaljawwad/pokedexcli/config"
	"github.com/mbilaljawwad/pokedexcli/internal/api"
	"github.com/savioxavier/termlink"
)

func mapBackCommand (conf *config.Config) error {
	url := conf.LocationUrl
	if (conf.PreviousUrl != nil) {
		url = *conf.PreviousUrl
	}

	locations := api.GetLocations(url)
	conf.SetPreviousUrl(locations.Previous)
	conf.SetNextUrl(locations.Next)
	
	fmt.Println()
	for _, location :=  range locations.Results {
		fmt.Println(termlink.Link(location.Name, location.Url))
	}
	fmt.Println()

	return nil
}