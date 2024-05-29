package asciiart

import (
	"fmt"
	"os"
	"strings"
)

func PrintAsciiArt(text string, asciiChars map[byte][]string, outputFile *os.File) {
	// Split the text by "\\n" to handle newlines
	splittxt := strings.Split(text, "\\n")
	countSpaces := 0
	for _, arg := range splittxt {
		if arg == "" {
			countSpaces++
			// Print newlines for empty lines except for the last one
			if countSpaces < len(splittxt) {
				fmt.Fprintln(outputFile)
			}
		} else {
			// Print ASCII art for non-empty lines
			PrintAscii(arg, asciiChars, outputFile)
		}
	}
}

func PrintAscii(text string, asciiChars map[byte][]string, outputFile *os.File) {
	// Replace escape sequences with their corresponding characters
	text = strings.NewReplacer(
		"\\r", "\r",
		"\\b", "\b",
		"\\t", "    ",
		"\\f", "\f",
		"\\a", "\a",
		"\\v", "\v",
	).Replace(text)

	// Check if any character is outside the ASCII range (32-127)
	for _, char := range text {
		if char > 127 || char < 32 {
			fmt.Printf("Error: Character %q is not accepted\n", char)
			return
		}
	}

	// Print each line of ASCII art for the given text
	for i := 0; i < 8; i++ {
		printLine(text, asciiChars, i, outputFile)
		fmt.Fprintln(outputFile)
	}
}

func printLine(text string, asciiChars map[byte][]string, line int, outputFile *os.File) {
	// Print each character's corresponding ASCII art for the given line
	for _, char := range text {
		if char == '\n' {
			// Print newline character
			fmt.Fprintln(outputFile)
		} else {
			// Print ASCII art for the character
			fmt.Fprint(outputFile, asciiChars[byte(char)][line])
		}
	}
}

