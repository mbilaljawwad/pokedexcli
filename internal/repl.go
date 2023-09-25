package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mbilaljawwad/pokedexcli/cmd"
)

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := cmd.GetCliCommands()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}

		cmdName := words[0]

		if command, ok := commands[cmdName]; ok {
			command.Callback()
		} else {
			fmt.Println("Command not found")
		}
	}
}

func cleanInput(input string) []string {
	output := strings.ToLower(input)
	words := strings.Fields(output)

	return words
}
