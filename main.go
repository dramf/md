package main

import (
	"io"
	"log"
	"os"

	"github.com/dramf/md/formatter"
)

const mdFile = "trivy-report.md"

func main() {
	inputData, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("error reading trivy output: %v\n", err)
	}

	err = formatter.Format(githubTpl, inputData, mdFile)
	if err != nil {
		log.Fatalf("error executing template: %v\n", err)
	}

	log.Printf("created Markdown file: %s\n", mdFile)
}
