package application

import (
	"context"
	"github.com/dnutiu/fluxa-cli/internal/domain"
)

type AccountRestorer interface {
	RestoreAccount(context.Context, domain.ID, domain.ID) (Record[domain.Account], error)
}

func RestoreAccount(ctx context.Context, repo AccountRestorer, entityID, accountID domain.ID) (Record[domain.Account], error) {
	if err := requireIDs(entityID, accountID); err != nil {
		return Record[domain.Account]{}, err
	}
	return repo.RestoreAccount(ctx, entityID, accountID)
}
