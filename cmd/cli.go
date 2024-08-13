package cmd

import (
	"flag"
	"fmt"
	"github.com/migrx-io/top10url/parser"
	"os"
)

// ParseArgs parses command-line arguments and
// returns the file path and the number of top URLs to extract.
func ParseArgs() (string, int) {

	// Define command-line flags
	filePath := flag.String("file", "", "The path to the input file")
	topN := flag.Int("top", 10, "Number of top URLs to output")

	// Parse the flags
	flag.Parse()

	// Ensure the file path is provided
	if *filePath == "" {
		fmt.Println("Usage: topurls -file=<file_path> [-top=<number_of_urls>]")
		os.Exit(1)
	}

	return *filePath, *topN
}

// ProcessFile reads the file, parses the data,
// and returns the top N URLs sorted by value.
func ProcessFile(filePath string, topN int) ([]string, error) {

	// Parse the file and get the URL-value pairs
	urlValues, err := parser.ParseFile(filePath, topN)
	if err != nil {
		return nil, err
	}

	return urlValues, nil
}
