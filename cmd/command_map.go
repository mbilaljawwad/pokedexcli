package cmd

import (
	"fmt"

	"github.com/mbilaljawwad/pokedexcli/api"
	"github.com/mbilaljawwad/pokedexcli/config"
	"github.com/savioxavier/termlink"
)

func mapCommand(conf *config.Config) error {
	url := conf.LocationUrl
	if (conf.NextUrl != nil) {
		url = *conf.NextUrl
	}
	locations := api.GetLocations(url)
	conf.SetNextUrl(locations.Next)
	conf.SetPreviousUrl(locations.Previous)
	
	fmt.Println()
	for _, location :=  range locations.Results {
		fmt.Println(termlink.Link(location.Name, location.Url))
	}
	fmt.Println()

	return nil
}