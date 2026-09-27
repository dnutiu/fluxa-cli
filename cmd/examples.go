package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var commandExamples = map[string]string{
	"fluxa":                       "fluxa auth status\nfluxa entities list",
	"fluxa auth":                  "fluxa auth status",
	"fluxa auth status":           "fluxa auth status",
	"fluxa config":                "fluxa config show\nfluxa config set-entity 7",
	"fluxa config show":           "fluxa config show",
	"fluxa config set-entity":     "fluxa config set-entity 7",
	"fluxa config set-url":        "fluxa config set-url http://127.0.0.1:3000",
	"fluxa entities":              "fluxa entities list",
	"fluxa entities list":         "fluxa entities list -o json",
	"fluxa accounts":              "fluxa accounts list -e 7",
	"fluxa accounts list":         "fluxa accounts list -e 7 --include-archived",
	"fluxa accounts show":         "fluxa accounts show 4 -e 7",
	"fluxa accounts add":          "fluxa accounts add -e 7 --name 'Everyday cash' --kind cash --currency RON",
	"fluxa accounts edit":         "fluxa accounts edit 4 -e 7 --name 'Travel savings' --opening-balance 100.00",
	"fluxa accounts archive":      "fluxa accounts archive 4 -e 7 --yes",
	"fluxa accounts restore":      "fluxa accounts restore 4 -e 7",
	"fluxa categories":            "fluxa categories list -e 7",
	"fluxa categories list":       "fluxa categories list -e 7",
	"fluxa transactions":          "fluxa transactions list -e 7 --kind expense",
	"fluxa transactions list":     "fluxa transactions list -e 7 -k expense -f 2026-09-01 -t 2026-09-30",
	"fluxa transactions show":     "fluxa transactions show 42 -e 7",
	"fluxa transactions add":      "fluxa transactions add -e 7 -g 12 -m 18.50 -d 2026-09-27 -n 'Coffee'",
	"fluxa transactions edit":     "fluxa transactions edit 42 -e 7 --description 'Team coffee' --exclude-from-analytics=false",
	"fluxa transactions delete":   "fluxa transactions delete 42 -e 7 --yes",
	"fluxa subscriptions":         "fluxa subscriptions list -e 7 --active=true",
	"fluxa subscriptions list":    "fluxa subscriptions list -e 7 --active=false",
	"fluxa subscriptions show":    "fluxa subscriptions show 12 -e 7",
	"fluxa subscriptions add":     "fluxa subscriptions add -e 7 -n 'Cloud storage' -m 12.50 -s 2026-09-27 -i monthly",
	"fluxa subscriptions edit":    "fluxa subscriptions edit 12 -e 7 --amount 14.50 --include-vat=true",
	"fluxa subscriptions pause":   "fluxa subscriptions pause 12 -e 7",
	"fluxa subscriptions resume":  "fluxa subscriptions resume 12 -e 7",
	"fluxa subscriptions delete":  "fluxa subscriptions delete 12 -e 7 --yes",
	"fluxa help":                  "fluxa help transactions add",
	"fluxa completion":            "fluxa completion zsh",
	"fluxa completion bash":       "source <(fluxa completion bash)",
	"fluxa completion fish":       "fluxa completion fish | source",
	"fluxa completion powershell": "fluxa completion powershell | Out-String | Invoke-Expression",
	"fluxa completion zsh":        "source <(fluxa completion zsh)",
}

func exampleRequested(cmd *cobra.Command) bool {
	show, err := cmd.Flags().GetBool("example")
	return err == nil && show
}

func installExamples(root *cobra.Command) {
	var visit func(*cobra.Command)
	visit = func(command *cobra.Command) {
		command.Example = commandExamples[command.CommandPath()]
		originalArgs, originalRunE, originalRun := command.Args, command.RunE, command.Run
		command.Args = func(cmd *cobra.Command, args []string) error {
			if exampleRequested(cmd) {
				return nil
			}
			if originalArgs != nil {
				return originalArgs(cmd, args)
			}
			return nil
		}
		command.RunE = func(cmd *cobra.Command, args []string) error {
			if exampleRequested(cmd) {
				_, err := fmt.Fprintln(cmd.OutOrStdout(), strings.TrimSpace(cmd.Example))
				return err
			}
			if originalRunE != nil {
				return originalRunE(cmd, args)
			}
			if originalRun != nil {
				originalRun(cmd, args)
				return nil
			}
			return cmd.Help()
		}
		command.Run = nil
		for _, child := range command.Commands() {
			visit(child)
		}
	}
	visit(root)
}
