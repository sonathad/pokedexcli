package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCmd struct {
	name        string
	description string
	cb          func()
}

func main() {
	Pokedex()
}

func Pokedex() {
	sc := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to Pokedex CLI!")
	for {
		ShowPrompt()

		if !sc.Scan() {
			break
		}

		currentCmd, ok := commandsMap()[sc.Text()]
		if !ok {
			fmt.Println("command not found")
			break
		}
		currentCmd.cb()
	}

	fmt.Println("Thank you for playing!")
}

func ShowPrompt() {
	fmt.Fprint(os.Stdout, "pokedex > ")
}

func commandsMap() map[string]cliCmd {
	helpCmd := func() {
		fmt.Println(`Usage:

help: Displays a help message
exit: Exit the Pokedex`)
	}

	exitCmd := func() {
		fmt.Println("exiting...")
		os.Exit(0)
	}

	nextMapCmd := func() {
		MapCmd("next")
	}

	prevMapCmd := func() {
		MapCmd("prev")
	}

	return map[string]cliCmd{
		"help": {
			name:        "help",
			description: "Displays a help message",
			cb:          helpCmd,
		},
		"exit": {
			name:        "exit",
			description: "Exits Pokedex",
			cb:          exitCmd,
		},
		"map": {
			name:        "map",
			description: "The map command displays the names of 20 location areas in the Pokemon world. Each subsequent call to map should display the next 20 locations",
			cb:          nextMapCmd,
		},
		"mapb": {
			name:        "mapb",
			description: "The mapb command displays the names of 20 location areas in the Pokemon world. Each subsequent call to mapb should display the previous 20 locations",
			cb:          prevMapCmd,
		},
	}
}
