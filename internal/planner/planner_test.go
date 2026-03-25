package planner

import (
	"testing"

	"github.com/Marlliton/gargs/internal/cli"
	"github.com/stretchr/testify/assert"
)

func Test_(t *testing.T) {
	tests := []struct {
		name  string
		cfg   cli.Config
		items []string
		want  [][]string
	}{
		{
			name:  "default puts everything in one batch",
			cfg:   cli.Config{},
			items: []string{"a", "b", "c"},
			want:  [][]string{{"a", "b", "c"}},
		},
		{
			name:  "max args splits into multiple batches",
			cfg:   cli.Config{MaxArgs: 2},
			items: []string{"a", "b", "c"},
			want:  [][]string{{"a", "b"}, {"c"}},
		},
		{
			name:  "replace token creates one batch per item",
			cfg:   cli.Config{ReplaceToken: "{}"},
			items: []string{"a", "b", "c"},
			want:  [][]string{{"a"}, {"b"}, {"c"}},
		},
		{
			name:  "max lines splits batches",
			cfg:   cli.Config{MaxArgs: 2},
			items: []string{"a", "b", "c", "d", "e"},
			want:  [][]string{{"a", "b"}, {"c", "d"}, {"e"}},
		},
		{
			name:  "empty input returns no batches",
			cfg:   cli.Config{},
			items: nil,
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildBatches(tt.cfg, tt.items)
			assert.Equal(t, tt.want, got)
		})
	}
}
