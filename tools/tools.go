package tools

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Function for reading input file (banner)
func ReadInput(filename string) string {
	buffer, err := os.ReadFile(filename)
	CheckError(err, "Error: failed to read infile: \""+filename+"\"!!")
	if len(buffer) == 0 {
		log.Fatalln("Error: infile is empty: \"" + filename + "\"!!")
	}
	return string(buffer)
}

// Function fo remove empty string
func RemoveEmptyStrings(slice []string) []string {
	var result []string
	for _, str := range slice {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}

// Check the template name
func CheckTemplate(arg string) string {
	var data string
	switch arg {
	case "standard":
		data = ReadInput("Templates/standard.txt")
	case "shadow":
		data = ReadInput("Templates/shadow.txt")
	case "thinkertoy":
		data = ReadInput("Templates/thinkertoy.txt")
	default:
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}
	return data
}

// Storing result in file output
func StoreResult(filename, content string) {
	if strings.HasPrefix(filename, "--output=") {
		filename = strings.TrimPrefix(filename, "--output=")
	}
	if filename == "main.go" {
		fmt.Println("Cannot store result in main.go")
		return
	}
	file, err := os.Create(filename)
	CheckError(err, "Error creating file: "+filename)
	defer file.Close()
	_, err = file.WriteString(content)
	CheckError(err, "Error writing to file: "+filename)
	fmt.Println("Result stored in", filename)
}

// function to take color
func TakeColor(color string) string {
	if strings.HasPrefix(color, "--color=") {
		color = strings.TrimPrefix(color, "--output=")
	}
	return color
}

// Ckeck if all input string new lines
func IsAllNl(result string) bool {
	for _, char := range result {
		if char != '\n' {
			return false
		}
	}
	return true
}

// Function for error
func CheckError(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
