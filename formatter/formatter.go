package formatter

import (
	"bytes"
	"encoding/json"
	"golang.org/x/xerrors"
	"log"
	"os"
	"text/template"
)

const mdFile = "trivy-report.md"

type Formatter struct {
	outputFile string
	template   *template.Template
}

func NewFormatter(options ...func(*Formatter)) (*Formatter, error) {
	tmpl, err := template.New("temp").Parse(githubTpl)
	if err != nil {
		return nil, xerrors.Errorf("error parsing template: %v\n", err)
	}

	formatter := &Formatter{
		outputFile: mdFile,
		template:   tmpl,
	}

	for _, opt := range options {
		opt(formatter)
	}
	return formatter, nil
}

func WithOutputFile(outputFile string) func(*Formatter) {
	return func(formatter *Formatter) {
		formatter.outputFile = outputFile
	}
}

func WithTemplate(t string) func(*Formatter) {
	return func(formatter *Formatter) {
		tmpl, err := template.New("temp").Parse(t)
		if err != nil {
			log.Println("error parsing template:", err)
			return
		}
		formatter.template = tmpl
	}
}
func (f *Formatter) Format(inputData []byte) error {
	var output Output
	if err := json.NewDecoder(bytes.NewReader(inputData)).Decode(&output); err != nil {
		return xerrors.Errorf("error decoding body: %v\n", err)
	}

	file, err := os.Create(f.outputFile)
	if err != nil {
		return xerrors.Errorf("error creating file: %v\n", err)
	}
	defer file.Close()

	if err := f.template.Execute(file, output.Results); err != nil {
		return xerrors.Errorf("error executing template: %v\n", err)
	}
	return nil
}
