package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) >= 2 {
		name := strings.Join(os.Args[1:], " ")
		fmt.Printf("Hello, %s!\n", name)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a name (or 'quit' to exit): ")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			fmt.Println("Hmmm..., I didn't quite catch the name")
			continue
		}

		if strings.ToLower(input) == "quit" || strings.ToLower(input) == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		fmt.Printf("Hello, %s!\n", input)

		fmt.Print("Continue? (y to continue, any other key to exit): ")
		if scanner.Scan() {
			response := strings.ToLower(strings.TrimSpace(scanner.Text()))
			if response != "y" && response != "yes" {
				fmt.Println("Goodbye!")
				break
			}
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}

}
