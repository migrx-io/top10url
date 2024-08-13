package main

import (
	"fmt"
	"log"
    "time"
    "os"

    "github.com/migrx-io/top10url/cmd"

)

func main() {

    logger := log.New(os.Stderr, "", log.LstdFlags)

	// Parse command-line arguments
	filePath, topN := cmd.ParseArgs()

    start := time.Now()

	// Read and parse the file
	urls, err := cmd.ProcessFile(filePath, topN)
	if err != nil {
		log.Fatalf("Error processing file: %v", err)
	}

	// Output the results
	for _, url := range urls {
		fmt.Println(url)
	}

    duration := time.Since(start)
    logger.Printf("Successfully processed file in %v\n", duration)

}
