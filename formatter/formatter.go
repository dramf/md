package formatter

import (
	"bytes"
	"encoding/json"
	"github.com/aquasecurity/trivy/pkg/types"
	"golang.org/x/xerrors"
	"os"
	"text/template"
)

const mdFile = "trivy-report.md"

type Formatter struct {
	outputFile string
	template   string
}

func NewFormatter(options ...func(*Formatter)) (*Formatter, error) {
	formatter := &Formatter{
		outputFile: mdFile,
		template:   githubTpl,
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
		formatter.template = t
	}
}
func (f *Formatter) Format(inputData []byte) error {
	var output types.Report
	if err := json.NewDecoder(bytes.NewReader(inputData)).Decode(&output); err != nil {
		return xerrors.Errorf("error decoding body: %v\n", err)
	}

	file, err := os.Create(f.outputFile)
	if err != nil {
		return xerrors.Errorf("error creating file: %v\n", err)
	}
	defer file.Close()

	tmpl, err := template.New("temp").Parse(f.template)
	if err != nil {
		return xerrors.Errorf("error parsing template: %v\n", err)
	}

	if err := tmpl.Execute(file, output.Results); err != nil {
		return xerrors.Errorf("error executing template: %v\n", err)
	}
	return nil
}
