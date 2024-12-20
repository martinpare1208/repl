package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/martinpare1208/pokedexcli/internal/client"
	"github.com/martinpare1208/pokedexcli/internal/commands"
	"github.com/martinpare1208/pokedexcli/internal/config"
)

func main() {
	
	pokeClient := client.NewClient(5*time.Second, time.Minute*5)

	fmt.Println("Welcome!")
	reader := bufio.NewScanner(os.Stdin)
	clientCfg := &config.Cfg{
		PokeClient: pokeClient,
	}

	// Create a REPL loop
	for {
		var command string
		var input string
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		input = ""
		
		if len(words) == 2 {
			input = words[1]
		}

		command = words[0]
		
	
		err := commands.ReadCommand(command, clientCfg, input)
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

