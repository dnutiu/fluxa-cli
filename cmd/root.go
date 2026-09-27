package cmd

import (
	"fmt"

	"github.com/dnutiu/fluxa-cli/internal/infrastructure/config"
	"github.com/spf13/cobra"
)

type app struct {
	settings   *config.Settings
	configFile string
}

func NewRootCommand() *cobra.Command {
	a := &app{settings: config.New()}
	root := &cobra.Command{
		Use: "fluxa", Short: "Manage Fluxa from the terminal",
		SilenceUsage: true, SilenceErrors: true,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			if exampleRequested(cmd) {
				return nil
			}
			return a.settings.Load(a.configFile)
		},
	}
	root.PersistentFlags().StringVarP(&a.configFile, "config", "C", "", "config file (default: user config directory/fluxa/config.yaml)")
	root.PersistentFlags().StringP("base-url", "u", "", "Fluxa server URL (or FLUXA_BASE_URL)")
	root.PersistentFlags().IntP("entity", "e", 0, "entity ID (or FLUXA_ENTITY)")
	root.PersistentFlags().StringP("output", "o", "", "output format: table or json (or FLUXA_OUTPUT)")
	root.PersistentFlags().BoolP("example", "X", false, "show an example for this command")
	if err := a.settings.BindFlags(root.PersistentFlags()); err != nil {
		panic(err)
	}
	root.AddCommand(a.configCommand(), a.authCommand(), a.entitiesCommand(), a.accountsCommand(), a.categoriesCommand(), a.transactionsCommand(), a.subscriptionsCommand())
	help := &cobra.Command{
		Use: "help [command]", Short: "Help about any command",
		RunE: func(cmd *cobra.Command, args []string) error {
			target, _, err := root.Find(args)
			if err != nil || target == nil {
				return fmt.Errorf("unknown help topic: %v", args)
			}
			return target.Help()
		},
	}
	root.SetHelpCommand(help)
	root.InitDefaultHelpCmd()
	root.InitDefaultCompletionCmd()
	installExamples(root)
	return root
}
