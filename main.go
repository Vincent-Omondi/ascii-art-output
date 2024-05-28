package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"ascii-art-output/asciiart"
)

func main() {
	outputFileName := flag.String("output", "", "Output file name (e.g., --output=banner.txt)")
	flag.Parse()

	if *outputFileName == "" || len(flag.Args()) < 1 {
		fmt.Println("Usage: go run . --output=<fileName.txt> something standard")
		return
	}

	bannerFile := "standard.txt"
	if len(flag.Args()) > 1 {
		bannerFile = flag.Args()[1] + ".txt"
	}

	asciiChars, err := asciiart.LoadAsciiChars(bannerFile)
	if err != nil {
		fmt.Println("Error loading ASCII characters:", err)
		return
	}

	text := strings.ReplaceAll(flag.Args()[0], "\\n", "\n")

	outputFile, err := os.Create(*outputFileName)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	asciiart.PrintAsciiArt(text, asciiChars, outputFile)
}
