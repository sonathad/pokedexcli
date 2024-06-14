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

		sc.Scan()
		text := sc.Text()
		if text == "" {
			break
		}
		currentCmd, ok := commandsMap()[text]
		if !ok {
			fmt.Println("command not found, thank you for playing!")
			break
		}
		currentCmd.cb()
	}
}

func ShowPrompt() {
	fmt.Fprint(os.Stdout, "pokedex > ")
}

func commandsMap() map[string]cliCmd {
	helpCmd := func() {
		fmt.Println("help is running")
	}

	exitCmd := func() {
		fmt.Println("exiting...")
		os.Exit(0)
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
	}
}
