package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/martinpare1208/pokedexcli/commands"
	
)

func main() {

	fmt.Println("Welcome!")
	reader := bufio.NewScanner(os.Stdin)

	// Create a REPL loop
	for {
		var command string
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		command = words[0]
	
		err := commands.ReadCommand(command)
		if err != nil {
			fmt.Println(err)
		}

	}
	
}

func cleanInput(command string) []string {
	output := strings.ToLower(command)
	words := strings.Fields(output)
	return words
}