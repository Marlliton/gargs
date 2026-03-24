// Package cli parse args
package cli

import (
	"strconv"
)

type Config struct {
	Command       string   // comando a ser executado (ex: "echo")
	FixedArgs     []string // argumentos fixos que sempre acompanham o comando
	MaxArgs       int      // -n: máximo de argumentos por execução (batch por quantidade de tokens)
	MaxLines      int      // -L: máximo de linhas por execução (batch por linhas, não tokens)
	NullDelimited bool     // -0: entrada separada por \0 em vez de whitespace
	ReplaceToken  string   // -I: placeholder a ser substituído (ativa modo 1 execução por item)
	PrintCommands bool     // -t: imprime o comando antes de executar (debug)
	NoRunIfEmpty  bool     // -r: não executa nada se não houver entrada
}

func Parse(args []string) (Config, error) {
	cfg := Config{}
	var foundCmd bool

	i := 0
	for i < len(args) {
		arg := args[i]

		if !foundCmd {
			flag, ok := isSupportedFlag(arg)

			if ok {
				switch flag.Name {
				case FlagMaxArgs:
					value, next, err := parseIntFlag(args, i, string(flag.Name))
					if err != nil {
						return cfg, err
					}
					cfg.MaxArgs = value

					i = next
					continue
				case FlagMaxLines:
					value, next, err := parseIntFlag(args, i, string(flag.Name))
					if err != nil {
						return cfg, err
					}

					cfg.MaxLines = value

					i = next
					continue
				case FlagNullDelimited:
					cfg.NullDelimited = true
					i++
					continue
				case FlagReplace:
					value, next, err := parseStringFlag(args, i, string(flag.Name))
					if err != nil {
						return cfg, err
					}

					cfg.ReplaceToken = value
					i = next
					continue
				case FlagPrint:
					cfg.PrintCommands = true
					i++
					continue
				case FlagNoRunIfEmpty:
					cfg.NoRunIfEmpty = true
					i++
					continue
				}
			}
		}

		if len(arg) > 0 && arg[0] == '-' {
			return cfg, UnknownFlagError{Flag: arg}
		}

		cfg.Command = arg
		foundCmd = true

		if len(args[i+1:]) > 0 {
			cfg.FixedArgs = args[i+1:]
		}

		break
	}

	return cfg, nil
}

func parseIntFlag(args []string, i int, flag string) (int, int, error) {
	targetPos := i + 1
	if targetPos >= len(args) {
		return 0, i, MissingValueError{Flag: flag}
	}
	n, err := strconv.Atoi(args[targetPos])
	if err != nil {
		return 0, i, InvalidValueError{Flag: flag}
	}

	return n, targetPos + 1, nil
}

func parseStringFlag(args []string, i int, flag string) (string, int, error) {
	var zero string
	targetPos := i + 1
	if targetPos >= len(args) {
		return zero, i, MissingValueError{Flag: flag}
	}

	return args[targetPos], targetPos + 1, nil
}
