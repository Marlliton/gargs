package input

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ReadInput(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		want          []string
		nullDelimited bool
	}{
		{
			name:  "whitespaces",
			input: "ana bia caio",
			want:  []string{"ana", "bia", "caio"},
		},
		{
			name:  "simple lines",
			input: "ana\nbia\ncaio\n",
			want:  []string{"ana", "bia", "caio"},
		},
		{
			name:          "null delimited",
			input:         "ana\x00bia\x00caio\x00",
			nullDelimited: true,
			want:          []string{"ana", "bia", "caio"},
		},
		{
			name:  "entrada vazia",
			input: "",
			want:  []string{},
		},
		{
			name:  "separa por whitespace",
			input: "ana maria\tjoao\ncaio",
			want:  []string{"ana", "maria", "joao", "caio"},
		},
		{
			name:          "separa por null byte",
			input:         "ana maria\x00joao\x00",
			nullDelimited: true,
			want:          []string{"ana maria", "joao"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.input)

			got, err := ReadInput(r, tt.nullDelimited)

			require.NoError(t, err)
			require.Equal(t, got, tt.want)
		})
	}
}
