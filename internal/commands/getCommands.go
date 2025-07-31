package commands

import "github.com/gnomedevreact/flat-world/internal/flat"

func GetCommands(characters *flat.Characters) (map[string]Command, error) {
	commands := make(map[string]Command)
	commands["help"] = Command{
		Handler:     HelpHandler(&commands),
		Description: `Displays a list of available commands and brief instructions for interacting with the Reality Show simulation.`,
	}
	commands["add"] = Command{
		Handler:     AddHandler(characters),
		Description: `Adds a new character to a reality Show simulation.`,
	}
	commands["all"] = Command{
		Handler:     GetAllHandler(characters),
		Description: "Show all characters",
	}
	commands["start"] = Command{
		Handler:     StartHandler(characters),
		Description: `Start a new simulation of a reality Show.`,
	}
	return commands, nil
}
