// Package planner build batches
package planner

import "github.com/Marlliton/gargs/internal/cli"

func BuildBatches(cfg cli.Config, items []string) [][]string {
	if len(items) == 0 {
		return nil
	}

	if cfg.ReplaceToken != "" {
		batches := make([][]string, 0, len(items))
		for _, it := range items {
			batches = append(batches, []string{it})
		}
		return batches
	}

	size := len(items)

	if cfg.MaxArgs > 0 {
		size = cfg.MaxArgs
	} else if cfg.MaxLines > 0 {
		size = cfg.MaxLines
	}

	batches := make([][]string, 0, (len(items)+size-1)/size)
	for i := 0; i < len(items); i += size {
		end := min(i+size, len(items))

		batches = append(batches, items[i:end])
	}

	return batches
}
