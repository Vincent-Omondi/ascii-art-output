package asciiart

import (
	"fmt"
	"os"
	"strings"
)

func PrintAsciiArt(text string, asciiChars map[byte][]string, outputFile *os.File) {
	splittxt := strings.Split(text, "\\n")
	countSpaces := 0
	for _, arg := range splittxt {
		if arg == "" {
			countSpaces++
			if countSpaces < len(splittxt) {
				fmt.Fprintln(outputFile)
			}
		} else {
			PrintAscii(arg, asciiChars, outputFile)
		}
	}
}

func PrintAscii(text string, asciiChars map[byte][]string, outputFile *os.File) {
	text = strings.NewReplacer(
		"\\r", "\r",
		"\\b", "\b",
		"\\t", "    ",
		"\\f", "\f",
		"\\a", "\a",
		"\\v", "\v",
	).Replace(text)
	for _, char := range text {
		if char > 127 || char < 32 {
			fmt.Printf("Error: Character %q is not accepted\n", char)
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
