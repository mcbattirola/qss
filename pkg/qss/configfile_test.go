package qss

import (
	"os"
	"testing"

	"github.com/mcbattirola/qss/pkg/require"
)

func TestParseConfigFile(t *testing.T) {
	tt := []struct {
		name     string
		content  string
		expected Config
	}{
		{
			name:     "empty",
			content:  "",
			expected: Config{},
		},
		{
			name: "with values",
			content: `
				font-size=24
				show-help=true
				save-path=/`,
			expected: Config{
				FontSize: 24,
				ShowHelp: true,
				FilePath: "/",
			},
		},
		{
			name: "ignores comments",
			content: `
				#font-size=24
				#show-help=true
				save-path=/`,
			expected: Config{
				FontSize: 0,
				ShowHelp: false,
				FilePath: "/",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tmpfile, err := os.CreateTemp("", "config")
			require.NoError(t, err)

			defer os.Remove(tmpfile.Name())

			_, err = tmpfile.Write([]byte(tc.content))
			require.NoError(t, err)

			_, err = tmpfile.Seek(0, 0)
			require.NoError(t, err)
			config := &Config{}

			parseConfigFile(tmpfile, config)
			require.Equal(t, &tc.expected, config)
		})
	}
}
