package commands

import (
	"fmt"
	"github.com/gnomedevreact/flat-world/internal/flat"
	"time"
)

func StartHandler(characters *flat.Characters) HandlerType {
	return func(args []string) error {
		if len(*characters) <= 1 {
			fmt.Println("Add at least 2 characters")
			return nil
		}

		tickChan := make(chan struct{}, 1)
		timer := time.NewTicker(30 * time.Second)
		dayNum := 0
		defer timer.Stop()

		fmt.Println("\033[34mSimulation launched\033[0m")

		go func() {
			for range timer.C {
				tickChan <- struct{}{}
			}
		}()

		tickChan <- struct{}{}

		for {
			fmt.Println("\n\033[35m===== Simulating... =====\033[0m")
			select {
			case <-tickChan:
				dayNum += 1
				fmt.Printf("\033[37m[Day %d]\033[0m\n", dayNum)

				for i := range *characters {
					message, err := flat.MakeDecision(&(*characters)[i], characters)
					if err != nil {
						return err
					}

					fmt.Println("\033[36m-----------------------------------\033[0m")
					fmt.Printf("\033[1;33m%s says:\033[0m\n", (*characters)[i].Name)
					fmt.Printf("  \033[32m\"%s\"\033[0m\n", *message)
				}

				fmt.Println("\033[36m-----------------------------------\033[0m")
			}
		}
	}
}
