package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"ascii-art-output/asciiart"
)

func main() {
	// Check if the command-line arguments are valid; if not, print usage instructions
	if len(os.Args) < 4 || !strings.HasPrefix(os.Args[1], "--output=") {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("Example: go run . --output=<fileName.txt> something standard")
		return
	}

	// Set up command-line flags
	outputFileName := flag.String("output", "", "Output file name (e.g., --output=banner.txt)")
	flag.Parse()

	// Check if output file name or text is not provided
	if *outputFileName == "" || len(flag.Args()) < 1 {
		fmt.Println("Usage: go run .--output=<outputFileName.txt> <textToConvert> <style>")
		return
	}

	// Set default banner style file
	bannerFile := "standard.txt"
	if len(flag.Args()) > 1 {
		bannerFile = flag.Args()[1] + ".txt" // Append ".txt" to style argument
	}

	// Load ASCII characters from specified style file
	asciiChars, err := asciiart.LoadAsciiChars(bannerFile)
	if err != nil {
		fmt.Println("Error loading ASCII characters:", err)
		return
	}

	// Replace newline characters in input text with "\\n" for proper handling
	text := strings.ReplaceAll(flag.Args()[0], "\n", "\\n")

	// Create the output file
	outputFile, err := os.Create(*outputFileName)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close() // Close the output file after function exits

	// Print ASCII art to the output file
	asciiart.PrintAsciiArt(text, asciiChars, outputFile)
}
