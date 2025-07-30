package commands

import (
	"errors"
	"fmt"
	"os"
)

func HelpHandler(cmds *map[string]Command) HandlerType {
	return func(args []string) error {
		if len(args) > 0 {
			return errors.New("no arguments expected")
		}
		for name, cmd := range *cmds {
			cmdName := fmt.Sprintf("|| \033[32m%s\033[0m", name)
			os.Stdout.WriteString(cmdName + "\n" + cmd.Description + "\n")
		}
		return nil
	}
}
