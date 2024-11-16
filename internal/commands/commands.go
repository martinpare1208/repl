package commands

import (
	"fmt"
	"os"

	"github.com/martinpare1208/pokedexcli/internal/config"
	"github.com/martinpare1208/pokedexcli/internal/pokeapi"

)



type CliCommand struct {
	Name string
	Desc string
	Callback func(cfg *config.Cfg, input string) error
}

type WrongCommandError struct {
	Command string
}


func getCommands() map[string]CliCommand{
return map[string]CliCommand{
	"help": {
		Name: "help",
		Desc: "display commands available",
		Callback: GetHelp,
	},
	"exit": {
		Name: "exit",
		Desc: "exit the program",
		Callback: ExitProgram,
	},
	"map": {
		Name: "map",
		Desc: "get information on location areas",
		Callback: getMap,
	},
	"mapb": {
		Name: "mapb",
		Desc: "go back a previous page",
		Callback: getMapB,
	},
	"explore": {
		Name: "explore",
		Desc: "explore an area",
		Callback: getPokemonData,
	},
	"catch": {
		Name: "catch",
		Desc: "catch a pokemon",
		Callback: catchPokemon,
	},
	"inspect": {
		Name: "inspect",
		Desc: "inspect a pokemon",
		Callback: inspectPokemon,
	},
}
}

func(c WrongCommandError) Error() string {
	return fmt.Sprintf("command: '%s' not found: ", c.Command)
}

func ExitProgram(cfg *config.Cfg, input string) (error) {
	fmt.Println("Exiting Program.")
	os.Exit(0)
	return nil
}

func GetHelp(cfg *config.Cfg, input string) (error) {

	fmt.Println("Commands below: ")
	fmt.Println()
	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n", c.Name, c.Desc)
	}
	return nil
}

func ReadCommand(command string, cfg *config.Cfg, input string) error {
	for c := range getCommands() {
		if command == c {
			getCommands()[command].Callback(cfg, input)
			return nil
		} 
	}
	return WrongCommandError{Command: command}
}

func getMap(cfg *config.Cfg, input string) (error) {
	err := pokeapi.GetLocations(cfg, cfg.NextUrl)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func getMapB(cfg *config.Cfg, input string) (error) {
	err := pokeapi.GetLocationsB(cfg, cfg.PrevUrl)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}

func getPokemonData(cfg *config.Cfg, location string) (error) {
	err := pokeapi.GetPokemonInArea(cfg, location)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func catchPokemon(cfg *config.Cfg, pokemonName string) (error) {
	err := pokeapi.CatchPokemon(cfg, pokemonName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}


func inspectPokemon(cfg *config.Cfg, pokemonName string) (error) {
	err := pokeapi.InspectPokemonInPokedex(cfg, pokemonName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
