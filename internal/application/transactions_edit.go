package application

import (
	"context"

	"github.com/dnutiu/fluxa-cli/internal/domain"
)

type TransactionUpdater interface {
	UpdateTransaction(context.Context, domain.ID, domain.ID, domain.TransactionInput) (Record[domain.Transaction], error)
}

func EditTransaction(ctx context.Context, repo TransactionUpdater, entityID, transactionID domain.ID, input domain.TransactionInput) (Record[domain.Transaction], error) {
	if err := requireIDs(entityID, transactionID); err != nil {
		return Record[domain.Transaction]{}, err
	}
	if err := input.ValidateUpdate(); err != nil {
		return Record[domain.Transaction]{}, err
	}
	return repo.UpdateTransaction(ctx, entityID, transactionID, input)
}
