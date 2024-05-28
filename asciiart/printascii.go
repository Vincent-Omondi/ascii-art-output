package asciiart

import (
	"fmt"
	"os"
	"strings"
)

func PrintAsciiArt(text string, asciiChars map[byte][]string, outputFile *os.File) {
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
