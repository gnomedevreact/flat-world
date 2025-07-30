package main

import (
	"bufio"
	"fmt"
	"github.com/gnomedevreact/flat-world/internal/commands"
	"github.com/gnomedevreact/flat-world/internal/flat"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	characters := flat.Characters{}
	commands, err := commands.GetCommands(&characters)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Reality Show CLI started. Type 'help'.")

	for {
		fmt.Print("> ")
		if !reader.Scan() {
			fmt.Println("Bye!")
			break
		}
		cmd := strings.TrimSpace(reader.Text())
		args := strings.Split(cmd, " ")

		err := commands[strings.ToLower(args[0])].Handler(args[1:])
		if err != nil {
			log.Fatal(err)
		}
	}
}
