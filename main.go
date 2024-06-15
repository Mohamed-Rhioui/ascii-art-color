package main

import (
	"fmt"
	"os"
	"strings"

	"ascii-art-output/programs"
)

func main() {
	args := os.Args[1:]
	// Protection
	if len(args) > 4 || len(args) == 0 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}
	outputToFile := false
	color := false
	// if (strings.HasPrefix(args[0], "--output=") && strings.HasPrefix(args[1], "--color=")) || (strings.HasPrefix(args[1], "--output=") && strings.HasPrefix(args[0], "--color=")) {
	// 	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
	// 	os.Exit(0)
	// }
	if strings.HasPrefix(args[0], "--output=") {
		outputToFile = true
	} else if strings.HasPrefix(args[0], "--color=") {
		color = true
	}

	// Check if the output file already exists
	if outputToFile {
		if len(args) == 3 || (len(args) == 2 && strings.HasPrefix(args[0], "--output=")) {
			programs.AsciiArt(true, false)
		} else if len(args) == 1 && strings.HasPrefix(args[0], "--output=") {
			programs.AsciiArt(false, false)
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
			os.Exit(0)
		}
	} else if color {
		if len(args) == 4 || len(args) == 3 || (len(args) == 2 && strings.HasPrefix(args[0], "--color=")) {
			programs.AsciiArt(false, true)
		} else if len(args) == 1 && strings.HasPrefix(args[0], "--color=") {
			programs.AsciiArt(false, false)
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
			os.Exit(0)
		}
	} else {
		if len(args) == 1 || len(args) == 2 {
			if strings.HasPrefix(args[0], "--output") || strings.HasPrefix(args[0], "--color") {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
				os.Exit(0)
			}
			programs.AsciiArt(false,false)
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
			os.Exit(0)
		}
	}
}
