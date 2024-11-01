package main
import (
	"fmt"
	"github.com/martinpare1208/pokedexcli/commands"
)

func main() {
	
	fmt.Println("Welcome!")
	// Create a REPL loop
	for {
		var command string
		
		fmt.Print("pokedex > ")
		fmt.Scanln(&command)
		err := commands.ReadCommand(command)
		if err == nil {
			fmt.Println("Command executed!")
		}

	}
	
}