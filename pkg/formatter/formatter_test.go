package formatter

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormatter(t *testing.T) {
	tests := []struct {
		name       string
		goldenPath string
		template   string
		jsonPath   string
	}{
		{
			name:       "happy all github",
			jsonPath:   "testdata/input/github/happy-all.json",
			goldenPath: "testdata/golden/github/golden-all.md",
		},
		{
			name:       "happy empty github",
			jsonPath:   "testdata/input/github/happy-empty.json",
			goldenPath: "testdata/golden/github/golden-empty.md",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cacheDir := t.TempDir()
			filePath := filepath.Join(cacheDir, mdFile)

			data, err := os.ReadFile(tt.jsonPath)
			require.NoError(t, err)

			f, err := NewFormatter(
				WithOutputFile(filePath))
			require.NoError(t, err)

			err = f.Format(data)
			require.NoError(t, err)

			actual, err := os.ReadFile(filePath)
			require.NoError(t, err)

			expected, err := os.ReadFile(tt.goldenPath)
			require.NoError(t, err)

			assert.Equal(t, string(expected), string(actual))
		})
	}
}
