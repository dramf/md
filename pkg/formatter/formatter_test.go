package formatter

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestFormatter(t *testing.T) {
	tests := []struct {
		name       string
		goldenPath string
		template   string
		jsonPath   string
	}{
		{
			name:       "happy all",
			jsonPath:   "testdata/input/happy-all.json",
			goldenPath: "testdata/golden/golden-all.md",
		},
		{
			name:       "happy empty",
			jsonPath:   "testdata/input/happy-empty.json",
			goldenPath: "testdata/golden/golden-empty.md",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cacheDir := t.TempDir()
			filePath := filepath.Join(cacheDir, mdFile)

			data, err := os.ReadFile(tt.jsonPath)
			assert.NoError(t, err)

			f, err := NewFormatter(
				WithOutputFile(filePath))
			assert.NoError(t, err)

			err = f.Format(data)
			assert.NoError(t, err)

			actual, err := os.ReadFile(filePath)
			assert.NoError(t, err)

			expected, err := os.ReadFile(tt.goldenPath)
			assert.NoError(t, err)

			assert.Equal(t, string(expected), string(actual))
		})
	}
}
