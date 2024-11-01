package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	inputData, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error reading from stdin: %v\n", err)
	}

	fmt.Printf("Trivy output data: %s\n", inputData)
}
