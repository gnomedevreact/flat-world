package commands

import (
	"fmt"
	"github.com/gnomedevreact/flat-world/internal/constants"
	"github.com/gnomedevreact/flat-world/internal/flat"
	"math/rand"
	"slices"
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
			Gender:             []string{"male", "female"}[rand.Intn(1)],
			Personality:        constants.Personalities[rand.Intn(len(constants.Personalities)-1)],
			Mood:               20,
			Hunger:             30,
			Patience:           80,
			Energy:             70,
			SocialNeed:         50,
			Interests:          _randomInterests(),
			Relationships:      make(map[string]int),
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

func _randomInterests() []string {
	interests := []string{}
	for len(interests) < 3 {
		randNum := rand.Intn(len(constants.Interests) - 1)
		randInterest := constants.Interests[randNum]

		if slices.Contains(interests, randInterest) {
			continue
		}

		interests = append(interests, randInterest)
	}
	return interests
}
