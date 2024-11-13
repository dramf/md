package main

import (
	"io"
	"log"
	"os"

	"github.com/dramf/md/pkg/formatter"
)

func main() {
	inputData, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("error reading trivy output: %v\n", err)
	}

	f, err := formatter.NewFormatter()
	if err != nil {
		log.Fatalf("error initializing formatter: %v\n", err)
	}

	err = f.Format(inputData)
	if err != nil {
		log.Fatalf("error executing template: %v\n", err)
	}
}
