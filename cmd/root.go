package cmd

import (
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
		PersistentPreRunE: func(_ *cobra.Command, _ []string) error { return a.settings.Load(a.configFile) },
	}
	root.PersistentFlags().StringVar(&a.configFile, "config", "", "config file (default: user config directory/fluxa/config.yaml)")
	root.PersistentFlags().String("base-url", "", "Fluxa server URL (or FLUXA_BASE_URL)")
	root.PersistentFlags().Int("entity", 0, "entity ID (or FLUXA_ENTITY)")
	root.PersistentFlags().String("output", "", "output format: table or json (or FLUXA_OUTPUT)")
	if err := a.settings.BindFlags(root.PersistentFlags()); err != nil {
		panic(err)
	}
	root.AddCommand(a.configCommand(), a.authCommand(), a.entitiesCommand(), a.accountsCommand(), a.categoriesCommand(), a.transactionsCommand(), a.subscriptionsCommand())
	root.InitDefaultCompletionCmd()
	return root
}
