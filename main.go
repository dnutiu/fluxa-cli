package main

import (
	"context"
	"fmt"
	"os"

	"fluxa-cli/cmd"
)

func main() {
	if err := cmd.NewRootCommand().ExecuteContext(context.Background()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
