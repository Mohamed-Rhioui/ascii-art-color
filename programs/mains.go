package programs

import (
	"fmt"
	"log"
	"os"
	"strings"

	"ascii-art-output/tools"
)

// // Function to draw ascii
func DrawAsciiArt(elements []string, input string,color string,toBeColored string) string {
	var result string
	lines := strings.Split(input, `\n`)
	for _, line := range lines {
		if line != "" {
			for j := 0; j < 8; j++ {
				for _, char := range line {
					if char < 32 || char > 126 {
						log.Fatalln("Error: please provide printable characters!!\nhelp: man ascii")
					} else {
						start := int(char-32)*8 + j
						result += elements[start]
					}
				}
				result += "\n"
			}
		} else {
			result += "\n"
		}
	}
	return result
}


// check if banner
func isTemplate(arg string) bool {
	var test bool
	switch arg {
	case "standard":
		test = true
	case "shadow":
		test = true
	case "thinkertoy":
		test = true
	default:
		test = false
	}

	return test
}

// Function to read input and template, and printing or storing result
func AsciiArt(isOutput, isColor bool) {
	args := os.Args[1:]
	var input string
	var color string
	var toBeColored string
	var template string

	if isOutput {
		input = args[1]
		if len(args) == 3 {
			template = args[2]
		} else {
			template = "standard"
		}
	} else if isColor {
		color = args[0]
		toBeColored = args[1]
		if len(args) == 4 {
			input = args[2]
			template = args[3]
		} else if len(args) == 3 {
			if isTemplate(args[2]) == true {
				input = args[1]
				template = args[2]
			} else {
				input = args[2]
			}
		}
	} else {
		input = args[0]
		if len(args) == 2 {
			template = args[1]
		} else {
			template = "standard"
		}
	}
	// Read template data
	data := tools.CheckTemplate(template)
	data = strings.ReplaceAll(data, "\r", "\n")
	elements := strings.Split(data, "\n")
	elements = tools.RemoveEmptyStrings(elements)

	// Draw ASCII art
	result := DrawAsciiArt(elements, input,color,toBeColored)

	// Handle additional new lines
	if tools.IsAllNl(result) {
		result = result[1:]
	}

	if isOutput {
		outputFile := args[0]
		tools.StoreResult(outputFile, result)
	} else {
		fmt.Print(result)
	}
}
