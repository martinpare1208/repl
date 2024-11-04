package commands

import(
	"fmt"
)

type CliCommand struct {
	Name string
	Desc string
	Callback func() error
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
}
}

func(c WrongCommandError) Error() string {
	return fmt.Sprintf("command: '%s' not found: ", c.Command)
}

func GetHelp() (error) {

	fmt.Println("Commands below: ")
	fmt.Println()
	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n", c.Name, c.Desc)
	}
	return nil
}

func ReadCommand(command string) error {
	for c := range getCommands() {
		if command == c {
			getCommands()[command].Callback()
			return nil
		} 
	}
	return WrongCommandError{Command: command}
}