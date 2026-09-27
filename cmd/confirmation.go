package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

func addYesFlag(cmd *cobra.Command) {
	cmd.Flags().BoolP("yes", "y", false, "confirm this action")
}

func requireYes(cmd *cobra.Command) error {
	yes, _ := cmd.Flags().GetBool("yes")
	if !yes {
		return errors.New("pass --yes or -y to confirm this action")
	}
	return nil
}
