package cmd

import (
	"fmt"

	"github.com/mbilaljawwad/pokedexcli/config"
)

var Cyan = "\033[36m"
var Reset = "\033[0m"

func helpCommand(conf *config.Config) error {
	fmt.Println("Welcome to the PokeDex CLI!")
	fmt.Println("Usage: ")
	fmt.Println()

	for _, command := range GetCliCommands(conf) {
		fmt.Printf("%s%s%s: %s\n", Cyan, command.Name, Reset, command.Description)
	}
	fmt.Println()

	return nil
}
