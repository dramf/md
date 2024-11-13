package formatter

import (
	"encoding/json"
	"os"

	"github.com/aquasecurity/trivy/pkg/types"
	"golang.org/x/xerrors"

	"github.com/dramf/md/pkg/renderer"
	"github.com/dramf/md/pkg/renderer/github"
)

const mdFile = "trivy-report.md"

type Formatter struct {
	outputFile string
	renderer   renderer.Renderer
}

func NewFormatter(options ...func(*Formatter)) (*Formatter, error) {
	formatter := &Formatter{
		outputFile: mdFile,
		renderer:   github.MarkdownGithubRenderer{},
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

func WithRenderer(r renderer.Renderer) func(*Formatter) {
	return func(formatter *Formatter) {
		formatter.renderer = r
	}
}
func (f *Formatter) Format(inputData []byte) error {
	var output types.Report
	err := json.Unmarshal(inputData, &output)
	if err != nil {
		return xerrors.Errorf("failed to unmarshal JSON: %w", err)
	}
	var result string
	if len(output.Results) == 0 {
		result = f.renderer.RenderEmptyTitle()
	}
	for _, r := range output.Results {
		if len(r.Vulnerabilities) != 0 || len(r.Misconfigurations) != 0 || len(r.Secrets) != 0 {
			result += f.renderer.RenderTitle(r.Target) +
				f.renderer.RenderVulnerabilities(r.Vulnerabilities) +
				f.renderer.RenderMisconfigurations(r.Misconfigurations) +
				f.renderer.RenderSecrets(r.Secrets)
		}
	}
	err = os.WriteFile(f.outputFile, []byte(result), 0600)
	if err != nil {
		return xerrors.Errorf("error creating file: %v\n", err)
	}

	return nil
}
