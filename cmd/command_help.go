package cmd

import "fmt"

var Cyan = "\033[36m"
var Reset = "\033[0m"

func helpCommand() error {
	fmt.Println("Welcome to the PokeDex CLI!")
	fmt.Println("Usage: ")
	fmt.Println()

	for _, command := range GetCliCommands() {
		fmt.Printf("%s%s%s: %s\n", Cyan, command.Name, Reset, command.Description)
	}
	fmt.Println()

	return nil
}
