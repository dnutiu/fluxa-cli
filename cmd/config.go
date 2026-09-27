package cmd

import (
	"encoding/json"
	"fmt"

	"fluxa-cli/internal/domain"
	"fluxa-cli/internal/infrastructure/fluxa"
	"github.com/spf13/cobra"
)

func (a *app) configCommand() *cobra.Command {
	parent := &cobra.Command{Use: "config", Short: "View or change local configuration"}
	parent.AddCommand(&cobra.Command{
		Use: "show", Short: "Show non-secret configuration", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return json.NewEncoder(cmd.OutOrStdout()).Encode(map[string]any{
				"base_url": a.settings.BaseURL(), "entity": a.settings.EntityID(),
				"output": a.settings.Output(), "file": a.settings.Path(),
			})
		},
	})
	parent.AddCommand(&cobra.Command{
		Use: "set-entity ID", Short: "Save a default entity ID", Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := domain.ParseID(args[0])
			if err != nil {
				return err
			}
			if err := a.settings.SaveEntity(id); err != nil {
				return err
			}
			_, err = fmt.Fprintln(cmd.OutOrStdout(), "Default entity saved.")
			return err
		},
	})
	parent.AddCommand(&cobra.Command{
		Use: "set-url URL", Short: "Save the Fluxa server URL", Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := fluxa.ValidateBaseURL(args[0]); err != nil {
				return err
			}
			if err := a.settings.SaveBaseURL(args[0]); err != nil {
				return err
			}
			_, err := fmt.Fprintln(cmd.OutOrStdout(), "Server URL saved.")
			return err
		},
	})
	return parent
}
