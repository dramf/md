package formatter

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

const mdFile = "trivy-report.md"

func TestFormatter(t *testing.T) {
	tests := []struct {
		name         string
		goldenPath   string
		templatePath string
		jsonPath     string
	}{
		{
			name:         "happy all",
			jsonPath:     "testdata/json/happy-all.json",
			templatePath: "testdata/template/github.tpl",
			goldenPath:   "testdata/golden/golden-all.md",
		},
		{
			name:         "happy empty",
			jsonPath:     "testdata/json/happy-empty.json",
			templatePath: "testdata/template/github.tpl",
			goldenPath:   "testdata/golden/golden-empty.md",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cacheDir := t.TempDir()
			filePath := filepath.Join(cacheDir, mdFile)

			data, err := os.ReadFile(tt.jsonPath)
			assert.NoError(t, err)

			template, err := os.ReadFile(tt.templatePath)
			assert.NoError(t, err)

			err = Format(string(template), data, filePath)
			assert.NoError(t, err)

			actual, err := os.ReadFile(filePath)
			assert.NoError(t, err)

			expected, err := os.ReadFile(tt.goldenPath)
			assert.NoError(t, err)

			assert.Equal(t, string(expected), string(actual))
		})
	}
}
