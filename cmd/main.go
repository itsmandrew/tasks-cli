package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func printHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  add <task description>   Add a new task")
	fmt.Println("  list                     List all tasks")
	fmt.Println("  exit                     Exit the application")
	fmt.Println("  help                     Show this help message")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("TODO CLI Application")
	fmt.Println("Type 'help' for a list of commands.")

	for {
		fmt.Print("> ") // Displays a prompt
		if !scanner.Scan() {
			break
		}

		// Reading and trimming input
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		// Split input into command and arguments
		args := strings.Fields(input)
		command := args[0]

		switch command {
		case "help":
			printHelp()

		case "exit":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Printf("Unknown command: %s\n", command)
			printHelp()
		}
		// TODO add more

	}

}
