package cmd

type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

func GetCliCommands() map[string]cliCommand {
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
	}
}
