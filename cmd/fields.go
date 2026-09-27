package cmd

import (
	"fmt"

	"github.com/dnutiu/fluxa-cli/internal/domain"
	"github.com/spf13/cobra"
)

func stringField(cmd *cobra.Command, name string) domain.Field[string] {
	if !cmd.Flags().Changed(name) {
		return domain.Field[string]{}
	}
	value, _ := cmd.Flags().GetString(name)
	return domain.With(value)
}

func boolField(cmd *cobra.Command, name string) domain.Field[bool] {
	if !cmd.Flags().Changed(name) {
		return domain.Field[bool]{}
	}
	value, _ := cmd.Flags().GetBool(name)
	return domain.With(value)
}

func idField(cmd *cobra.Command, name string) domain.Field[domain.ID] {
	if !cmd.Flags().Changed(name) {
		return domain.Field[domain.ID]{}
	}
	value, _ := cmd.Flags().GetInt64(name)
	return domain.With(domain.ID(value))
}

func nullableIDField(cmd *cobra.Command, name, clearName string) (domain.Field[domain.ID], error) {
	clear, _ := cmd.Flags().GetBool(clearName)
	if clear && cmd.Flags().Changed(name) {
		return domain.Field[domain.ID]{}, fmt.Errorf("--%s and --%s cannot be used together", name, clearName)
	}
	if clear {
		return domain.Cleared[domain.ID](), nil
	}
	return idField(cmd, name), nil
}

func nullableStringField(cmd *cobra.Command, name, clearName string) (domain.Field[string], error) {
	clear, _ := cmd.Flags().GetBool(clearName)
	if clear && cmd.Flags().Changed(name) {
		return domain.Field[string]{}, fmt.Errorf("--%s and --%s cannot be used together", name, clearName)
	}
	if clear {
		return domain.Cleared[string](), nil
	}
	return stringField(cmd, name), nil
}
