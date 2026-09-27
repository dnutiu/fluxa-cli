package application

import (
	"context"
	"fluxa-cli/internal/domain"
)

type AccountLister interface {
	ListAccounts(context.Context, domain.ID, domain.AccountFilter) (Collection[domain.Account], error)
}

func ListAccounts(ctx context.Context, repo AccountLister, entityID domain.ID, filter domain.AccountFilter) (Collection[domain.Account], error) {
	if err := requireIDs(entityID); err != nil {
		return Collection[domain.Account]{}, err
	}
	return repo.ListAccounts(ctx, entityID, filter)
}
