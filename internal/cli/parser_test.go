package cli

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Parser(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    Config
		wantErr error
	}{
		{
			name: "simple command",
			args: []string{"echo"},
			want: Config{Command: "echo"},
		},
		{
			name: "command with fixed args",
			args: []string{"echo", "a", "b"},
			want: Config{Command: "echo", FixedArgs: []string{"a", "b"}},
		},
		{
			name: "flag -n",
			args: []string{"-n", "3", "echo"},
			want: Config{Command: "echo", MaxArgs: 3},
		},
		{
			name: "flag -L",
			args: []string{"-L", "2", "echo"},
			want: Config{Command: "echo", MaxLines: 2},
		},
		{
			name: "flag -0",
			args: []string{"-0", "echo"},
			want: Config{Command: "echo", NullDelimited: true},
		},
		{
			name: "flag -t",
			args: []string{"-t", "echo"},
			want: Config{Command: "echo", PrintCommands: true},
		},
		{
			name: "flag -r",
			args: []string{"-r", "echo"},
			want: Config{Command: "echo", NoRunIfEmpty: true},
		},
		{
			name: "flag -I",
			args: []string{"-I", "{}", "echo"},
			want: Config{Command: "echo", ReplaceToken: "{}"},
		},
		{
			name: "flags combinadas",
			args: []string{"-0", "-n", "3", "echo", "x"},
			want: Config{
				Command:       "echo",
				FixedArgs:     []string{"x"},
				MaxArgs:       3,
				NullDelimited: true,
			},
		},
		{
			name:    "flag desconhecida",
			args:    []string{"-x", "echo"},
			wantErr: ErrUnknownFlag,
		},
		{
			name:    "faltando valor em -n",
			args:    []string{"-n"},
			wantErr: ErrMissingValue,
		},
		{
			name:    "valor invalido em -n",
			args:    []string{"-n", "abc", "echo"},
			wantErr: ErrInvalidValue,
		},
		{
			name: "flag depois do comando vira arg fixo",
			args: []string{"echo", "-n", "3"},
			want: Config{
				Command:   "echo",
				FixedArgs: []string{"-n", "3"},
			},
		},
		{
			name: "sem argumentos",
			args: nil,
			want: Config{},
		},
		{
			name: "so flag sem comando",
			args: []string{"-0"},
			want: Config{
				NullDelimited: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args)

			if tt.wantErr != nil {
				require.Error(t, err)
				require.True(t, errors.Is(err, tt.wantErr))
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
