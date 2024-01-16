package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
)

var (
	tried_passwords map[string]int
)

func load() {
	content, err := os.ReadFile("tried-passwords.txt")
	if err != nil {
		fmt.Println("No tried-passwords.txt file found.")
	}
	lines := strings.Split(string(content), "\n")

	tried_passwords = make(map[string]int)

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}

		tried_passwords[lines[i]] = 1
	}
}

func save() {
	file, err := os.Create("tried-passwords.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	firstIteration := true
	for key := range tried_passwords {
		// "\n" delimitter on all elements (except the last)
		if !firstIteration {
			_, err = file.WriteString("\n")
			if err != nil {
				log.Fatal(err)
			}
		}

		// Write the key as usual.
		_, err = file.WriteString(key)
		if err != nil {
			log.Fatal(err)
		}
		firstIteration = false
	}

	fmt.Println("Successfully saved.")
}

func main() {
	defer save()
	load()

	color.Set(color.BgWhite)
	color.Set(color.FgGreen)

	fmt.Printf(" Welcome to Tried Passwords CLI ")
	color.Set(color.Reset)
	fmt.Print("\n")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter a password (or '!exit' to quit): ")
		scanner.Scan()

		input := scanner.Text()

		if input == "!exit" {
			break
		}

		if _, ok := tried_passwords[input]; ok {
			color.Set(color.FgRed)
			fmt.Println("Already tried that.")
			color.Set(color.Reset)
		} else {
			tried_passwords[input] = 1
			color.Set(color.FgGreen)
			fmt.Println("New Password! Just added to \"tried passwords\" memory.")
			color.Set(color.Reset)
			save()
		}

		fmt.Println()
	}

	fmt.Println("Exiting...")
}
