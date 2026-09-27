package cmd

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestEveryCommandPrintsExampleWithoutCredentialsOrArguments(t *testing.T) {
	t.Setenv("FLUXA_API_KEY", "")
	var visit func(*cobra.Command)
	visit = func(cmd *cobra.Command) {
		path := cmd.CommandPath()
		if strings.TrimSpace(cmd.Example) == "" {
			t.Errorf("%s has no example", path)
		}
		args := append(strings.Fields(path)[1:], "--example")
		output, err := runCommand(t, args...)
		if err != nil {
			t.Errorf("%s --example: %v", path, err)
		} else if !strings.Contains(output, "fluxa ") {
			t.Errorf("%s --example output = %q", path, output)
		}
		for _, child := range cmd.Commands() {
			visit(child)
		}
	}
	visit(NewRootCommand())
}
