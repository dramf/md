package formatter

import (
	"bytes"
	"encoding/json"
	"golang.org/x/xerrors"
	"log"
	"os"
	"text/template"
)

func Format(tpl string, inputData []byte, outputFile string) error {
	var output Output
	if err := json.NewDecoder(bytes.NewReader(inputData)).Decode(&output); err != nil {
		log.Fatalf("error decoding body: %v\n", err)
	}

	tmpl, err := template.New("temp").Parse(tpl)
	if err != nil {
		return xerrors.Errorf("error parsing template: %v\n", err)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return xerrors.Errorf("error creating file: %v\n", err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, output.Results); err != nil {
		return xerrors.Errorf("error executing template: %v\n", err)
	}
	return nil
}
