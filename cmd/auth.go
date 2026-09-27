package cmd

import (
	"fmt"

	"fluxa-cli/internal/application"
	"fluxa-cli/internal/presentation"
	"github.com/spf13/cobra"
)

func (a *app) authCommand() *cobra.Command {
	parent := &cobra.Command{Use: "auth", Short: "Check API-key access"}
	parent.AddCommand(&cobra.Command{
		Use: "status", Short: "Validate FLUXA_API_KEY", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			repo, err := a.repository()
			if err != nil {
				return err
			}
			result, err := application.CheckAuthentication(cmd.Context(), repo)
			if err != nil {
				return err
			}
			if a.settings.Output() == "json" {
				return a.print(cmd, result, presentation.Entities)
			}
			_, err = fmt.Fprintf(cmd.OutOrStdout(), "Authenticated; %d entities available.\n", len(result.Data))
			return err
		},
	})
	return parent
}
