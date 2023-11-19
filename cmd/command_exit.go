package cmd

import (
	"os"

	"github.com/mbilaljawwad/pokedexcli/config"
)

func exitCommand(conf *config.Config) error {
	os.Exit(0)
	return nil
}
