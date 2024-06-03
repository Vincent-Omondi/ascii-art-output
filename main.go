package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"ascii-art-output/asciiart"
)

func main() {
	// Set up command-line flags
	outputFileName := flag.String("output", "", "Output file name (e.g., --output=banner.txt)")
	flag.Usage = func() {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("Example: go run . --output=<fileName.txt> something standard")
	}
	if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "--") {
		if !strings.HasPrefix(os.Args[1], "--output=") || len(os.Args) < 3 {
			flag.Usage()
			return
		}
	}
	flag.Parse()

	if len(os.Args) == 2 && strings.HasPrefix(os.Args[1], "--output") {
		flag.Usage()
		return
	}

	var textToConvert string
	var bannerFile string

	// Determine if the program is invoked with flags or arguments
	if *outputFileName == "" {
		// Handle the case where the program is run with arguments
		if len(os.Args) < 2 {
			flag.Usage()
			return
		}
		textToConvert = os.Args[1]
		if len(os.Args) > 2 {
			bannerFile = os.Args[2] + ".txt"
		} else {
			bannerFile = "standard.txt"
		}
	} else {
		// Handle the case where the program is run with flags
		if len(flag.Args()) < 1 {
			flag.Usage()
			return
		}
		textToConvert = flag.Args()[0]
		if len(flag.Args()) > 1 {
			bannerFile = flag.Args()[1] + ".txt"
		} else {
			bannerFile = "standard.txt"
		}
	}

	if *outputFileName != "" && !strings.HasSuffix(*outputFileName, ".txt") {
		fmt.Println("Error: Output file must have a .txt extension")
		return
	}

	// Load ASCII characters from the specified style file
	asciiChars, err := asciiart.LoadAsciiChars(bannerFile)
	if err != nil {
		fmt.Println("Error loading ASCII characters:", err)
		return
	}

	// Replace newline characters in input text with "\\n" for proper handling
	textToConvert = strings.ReplaceAll(textToConvert, "\n", "\\n")

	// Determine the output file
	var outputFile *os.File
	if *outputFileName == "" {
		outputFile, err = os.Create("banner.txt")
	} else {
		outputFile, err = os.Create(*outputFileName)
	}
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close() // Close the output file after function exits

	// Print ASCII art to the output file
	asciiart.PrintAsciiArt(textToConvert, asciiChars, outputFile)
}
