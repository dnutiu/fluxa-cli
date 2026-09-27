package application

import (
	"context"
	"fluxa-cli/internal/domain"
)

type AccountGetter interface {
	GetAccount(context.Context, domain.ID, domain.ID) (Record[domain.Account], error)
}

func ShowAccount(ctx context.Context, repo AccountGetter, entityID, accountID domain.ID) (Record[domain.Account], error) {
	if err := requireIDs(entityID, accountID); err != nil {
		return Record[domain.Account]{}, err
	}
	return repo.GetAccount(ctx, entityID, accountID)
}
