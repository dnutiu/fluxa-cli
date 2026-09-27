package application

import (
	"context"

	"github.com/dnutiu/fluxa-cli/internal/domain"
)

type AccountUpdater interface {
	UpdateAccount(context.Context, domain.ID, domain.ID, domain.AccountInput) (Record[domain.Account], error)
}

func EditAccount(ctx context.Context, repo AccountUpdater, entityID, accountID domain.ID, input domain.AccountInput) (Record[domain.Account], error) {
	if err := requireIDs(entityID, accountID); err != nil {
		return Record[domain.Account]{}, err
	}
	if err := input.ValidateUpdate(); err != nil {
		return Record[domain.Account]{}, err
	}
	return repo.UpdateAccount(ctx, entityID, accountID, input)
}
