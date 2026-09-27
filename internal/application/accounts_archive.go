package application

import (
	"context"
	"fluxa-cli/internal/domain"
)

type AccountArchiver interface {
	ArchiveAccount(context.Context, domain.ID, domain.ID) error
}

func ArchiveAccount(ctx context.Context, repo AccountArchiver, entityID, accountID domain.ID) error {
	if err := requireIDs(entityID, accountID); err != nil {
		return err
	}
	return repo.ArchiveAccount(ctx, entityID, accountID)
}
