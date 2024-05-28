package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	outputFileName := flag.String("output", "", "Output file name (e.g., --output=banner.txt)")
	flag.Usage = func() {
		fmt.Println("EX: go run . --output=<fileName.txt> something standard")
		// flag.PrintDefaults()
	}
	flag.Parse()

	if *outputFileName == "" {
		flag.Usage()
		return
	}

	args := flag.Args()

	if len(args) < 1 {
		flag.Usage()
		return
	}

	bannerFile := "standard.txt" // Default banner file
	if len(args) > 1 {
		bannerFile = args[1] + ".txt"
	}

	asciiChars, err := LoadAsciiChars(bannerFile)
	if err != nil {
		fmt.Println("Error loading ASCII characters:", err)
		return
	}

	text := strings.Join(args[:1], " ")
	text = strings.ReplaceAll(text, "\\n", "\n")

	outputFile, err := os.Create(*outputFileName)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	printAsciiArt(text, asciiChars, outputFile)
}

func printAsciiArt(text string, asciiChars map[byte][]string, outputFile *os.File) {
	text = strings.ReplaceAll(text, "\\r", "\r")
	text = strings.ReplaceAll(text, "\\b", "\b")
	text = strings.ReplaceAll(text, "\\t", "    ")
	text = strings.ReplaceAll(text, "\\f", "\f")
	text = strings.ReplaceAll(text, "\\a", "\a")
	text = strings.ReplaceAll(text, "\\v", "\v")
	for _, char := range text {
		if char > 127 || char < 32 {
			fmt.Fprintf(outputFile, "Error: Character %q is not accepted\n", char)
			return
		}
	}
	for i := 0; i < 8; i++ {
		printLine(text, asciiChars, i, outputFile)
		fmt.Fprintln(outputFile)
	}
}

func printLine(text string, asciiChars map[byte][]string, line int, outputFile *os.File) {
	for _, char := range text {
		if char == '\n' {
			fmt.Fprintln(outputFile)
		} else {
			fmt.Fprint(outputFile, asciiChars[byte(char)][line])
		}
	}
}

func LoadAsciiChars(filename string) (map[byte][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("error: file '%s' not found", filename)
		} else {
			return nil, fmt.Errorf("error opening file: %w", err)
		}
	}
	defer file.Close()

	asciiChars := make(map[byte][]string)

	scanner := bufio.NewScanner(file)

	currentChar := byte(' ')
	count := 0
	scanner.Scan() // Skip the first line
	for scanner.Scan() {
		line := scanner.Text()
		if count != 8 {
			asciiChars[currentChar] = append(asciiChars[currentChar], line)
			count++
		} else {
			currentChar++
			count = 0
		}
	}
	if len(asciiChars) == 0 {
		return nil, fmt.Errorf("error: file '%s' is empty", filename)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return asciiChars, nil
}
