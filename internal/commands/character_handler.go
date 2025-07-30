package commands

import (
	"fmt"
	"github.com/gnomedevreact/flat-world/internal/flat"
)

func AddHandler(characters *flat.Characters) HandlerType {
	return func(args []string) error {
		fmt.Println(args[0])
		if len(args) != 1 {
			fmt.Println("Usage: add <character name>")
			return nil
		}

		*characters = append(*characters, flat.Character{
			Name:               args[0],
			Personality:        "peaceful",
			Mood:               20,
			Hunger:             30,
			Patience:           80,
			Energy:             70,
			SocialNeed:         50,
			Interests:          []string{"music", "philosophy"},
			Relationships:      make(map[string]int),
			CurrentAction:      "idle",
			TicksSinceLastMeal: 3,
		})
		fmt.Println(fmt.Sprintf("\033[36mCharacter %s was added\033[0m", args[0]))

		return nil
	}
}

func GetAllHandler(characters *flat.Characters) HandlerType {
	return func(args []string) error {
		for _, character := range *characters {
			fmt.Println("\033[33m" + character.Name + "\033[0m")
		}
		return nil
	}
}
