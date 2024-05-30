package asciiart

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestLoadAsciiChars(t *testing.T) {
	// Create an empty standard.txt file for testing
	if err := ioutil.WriteFile("standard.txt", nil, 0o644); err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
	defer os.Remove("standard.txt")

	tests := []struct {
		name          string
		filename      string
		expectedMap   map[byte][]string
		expectedError string
	}{
		{
			name:          "Non-existent file",
			filename:      "nonexistent.txt",
			expectedMap:   nil,
			expectedError: "file 'nonexistent.txt' not found",
		},
		{
			name:          "Empty file",
			filename:      "standard.txt",
			expectedMap:   nil,
			expectedError: "error: file 'standard.txt' is empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asciiChars, err := LoadAsciiChars(tt.filename)

			if err != nil && err.Error() != tt.expectedError {
				t.Errorf("Got error '%v', expected '%v'", err.Error(), tt.expectedError)
			}

			if !reflect.DeepEqual(asciiChars, tt.expectedMap) {
				t.Errorf("Got map '%v', expected '%v'", asciiChars, tt.expectedMap)
			}
		})
	}
}
