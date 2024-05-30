package asciiart

import (
	"bufio"
	"fmt"
	"os"
)

// LoadAsciiChars loads ASCII characters from a file and returns a map
func LoadAsciiChars(filename string) (map[byte][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		// Check if the file doesn't exist
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("file '%s' not found", filename)
			// Return an error if there was a problem opening the file
		} else {
			return nil, fmt.Errorf("opening file: %w", err)
		}
	}
	defer file.Close()

	// Create a scanner to read from the file
	asciiChars := make(map[byte][]string)

	scanner := bufio.NewScanner(file)

	currentChar := byte(' ')
	count := 0

	// Skip the first line of the file
	scanner.Scan()

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
