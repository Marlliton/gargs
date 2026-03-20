package main

import (
	"fmt"
	"os"

	"github.com/Marlliton/gargs/internal/cli"
)

func main() {
	args := os.Args[1:]
	cfg, err := cli.Parse(args)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("args", cfg)
}
