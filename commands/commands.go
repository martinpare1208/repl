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

var CommandsList = map[string]CliCommand{

	"help": {
		Name: "help",
		Desc: "display commands available",
		Callback: getHelp,
	},

}

func(c WrongCommandError) Error() string {
	return fmt.Sprintf("command: '%s' not found: ", c.Command)
}

func getHelp() (error) {

	fmt.Println("Help successful")
	return nil
}

func ReadCommand(command string) error {
	for c := range CommandsList {
		if command == c {
			err := CommandsList[command].Callback()
			if err == nil {
				return nil
			}

		}
	}
	return WrongCommandError{Command: command}
}