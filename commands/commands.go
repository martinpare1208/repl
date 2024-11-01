package commands

import(
	"fmt"
)

type cliCommand struct {
	name string
	desc string
	callback func() error
}

var CommandsList = map[string]cliCommand{
	"help": {
		name: "help",
		desc: "display commands available",
		callback: getHelp,
	},

}

func getHelp() error {
	fmt.Println("Help successful")
	return nil
}