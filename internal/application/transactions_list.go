package application

import (
	"context"
	"github.com/dnutiu/fluxa-cli/internal/domain"
)

type TransactionLister interface {
	ListTransactions(context.Context, domain.ID, domain.TransactionFilter) (Collection[domain.Transaction], error)
}

func ListTransactions(ctx context.Context, repo TransactionLister, entityID domain.ID, filter domain.TransactionFilter) (Collection[domain.Transaction], error) {
	if err := requireIDs(entityID); err != nil {
		return Collection[domain.Transaction]{}, err
	}
	if err := filter.Validate(); err != nil {
		return Collection[domain.Transaction]{}, err
	}
	return repo.ListTransactions(ctx, entityID, filter)
}
