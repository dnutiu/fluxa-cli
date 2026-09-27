package cmd

import (
	"fluxa-cli/internal/application"
	"fluxa-cli/internal/presentation"
	"github.com/spf13/cobra"
)

func (a *app) categoriesCommand() *cobra.Command {
	parent := &cobra.Command{Use: "categories", Short: "List enabled categories"}
	parent.AddCommand(&cobra.Command{
		Use: "list", Short: "List enabled categories", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			entityID, err := a.entityID()
			if err != nil {
				return err
			}
			repo, err := a.repository()
			if err != nil {
				return err
			}
			result, err := application.ListCategories(cmd.Context(), repo, entityID)
			if err != nil {
				return err
			}
			return a.print(cmd, result, presentation.Categories)
		},
	})
	return parent
}
