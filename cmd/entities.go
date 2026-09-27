package cmd

import (
	"fluxa-cli/internal/application"
	"fluxa-cli/internal/presentation"
	"github.com/spf13/cobra"
)

func (a *app) entitiesCommand() *cobra.Command {
	parent := &cobra.Command{Use: "entities", Short: "Find your Fluxa entities"}
	parent.AddCommand(&cobra.Command{
		Use: "list", Short: "List entities", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			repo, err := a.repository()
			if err != nil {
				return err
			}
			result, err := application.ListEntities(cmd.Context(), repo)
			if err != nil {
				return err
			}
			return a.print(cmd, result, presentation.Entities)
		},
	})
	return parent
}
