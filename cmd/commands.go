package cmd

import "github.com/mbilaljawwad/pokedexcli/config"

type cliCommand struct {
	Name        string
	Description string
	Callback    func(conf *config.Config) error
}

func GetCliCommands(conf *config.Config) map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    helpCommand,
		},
		"exit": {
			Name:        "exit",
			Description: "Exits the pokedexcli",
			Callback:    exitCommand,
		},
		"map": {
			Name:        "map",
			Description: "Shows forward set of maps of the pokemon world by calling Locations API",
			Callback: mapCommand,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Shows previous set of maps of the pokemon world by calling Locations API",
			Callback: mapBackCommand,
		},
	}
}
