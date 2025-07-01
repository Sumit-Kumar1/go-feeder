package main

import (
	"context"
	"fmt"
	"os"

	"feeder/cmd"
)

func main() {
	ctx := context.Background()

	if err := cmd.Run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
