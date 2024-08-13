package main

import (
	"fmt"
	"github.com/migrx-io/top10url/cmd"
	"log"
)

func main() {

	// Parse command-line arguments
	filePath, topN := cmd.ParseArgs()

	// Read and parse the file
	urls, err := cmd.ProcessFile(filePath, topN)
	if err != nil {
		log.Fatalf("Error processing file: %v", err)
	}

	// Output the results
	for _, url := range urls {
		fmt.Println(url)
	}

}
