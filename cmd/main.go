package main

import (
	"bufio"
	"fmt"
	"github.com/gnomedevreact/flat-world/internal/commands"
	"github.com/gnomedevreact/flat-world/internal/flat"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	reader := bufio.NewScanner(os.Stdin)
	characters := flat.Characters{}
	commands, err := commands.GetCommands(&characters)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(`
██████╗ ███████╗ █████╗ ██╗     ██╗████████╗██╗   ██╗     ███████╗██╗  ██╗ ██████╗ ██╗    ██╗
██╔══██╗██╔════╝██╔══██╗██║     ██║╚══██╔══╝╚██╗ ██╔╝     ██╔════╝██║  ██║██╔═══██╗██║    ██║
██████╔╝█████╗  ███████║██║     ██║   ██║    ╚████╔╝█████╗███████╗███████║██║   ██║██║ █╗ ██║
██╔═══╝ ██╔══╝  ██╔══██║██║     ██║   ██║     ╚██╔╝ ╚════╝╚════██║██╔══██║██║   ██║██║███╗██║
██║     ███████╗██║  ██║███████╗██║   ██║      ██║        ███████║██║  ██║╚██████╔╝╚███╔███╔╝
╚═╝     ╚══════╝╚═╝  ╚═╝╚══════╝╚═╝   ╚═╝      ╚═╝        ╚══════╝╚═╝  ╚═╝ ╚═════╝  ╚══╝╚══╝ 
                             WELCOME TO THE REALITY SIMULATION
                           'help' - list all available commands
`)

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
