package application

import (
	"errors"

	"github.com/dnutiu/fluxa-cli/internal/domain"
)

func requireIDs(ids ...domain.ID) error {
	for _, id := range ids {
		if !id.Valid() {
			return errors.New("IDs must be positive integers")
		}
	}
	return nil
}
