package cmd

import (
	"errors"
	"os"

	"fluxa-cli/internal/domain"
	"fluxa-cli/internal/infrastructure/fluxa"
	"fluxa-cli/internal/presentation"
	"github.com/spf13/cobra"
)

func (a *app) repository() (*fluxa.Repository, error) {
	return fluxa.NewRepository(a.settings.BaseURL(), os.Getenv("FLUXA_API_KEY"), nil)
}

func (a *app) entityID() (domain.ID, error) {
	id := a.settings.EntityID()
	if !id.Valid() {
		return 0, errors.New("select an entity with --entity, FLUXA_ENTITY, or 'fluxa config set-entity ID'")
	}
	return id, nil
}

func (a *app) print(cmd *cobra.Command, result any, columns []presentation.Column) error {
	return presentation.Print(cmd.OutOrStdout(), result, a.settings.Output(), columns)
}
