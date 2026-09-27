package application

import (
	"context"
	"fluxa-cli/internal/domain"
)

type TransactionUpdater interface {
	UpdateTransaction(context.Context, domain.ID, domain.ID, []byte) (Record[domain.Transaction], error)
}

func EditTransaction(ctx context.Context, repo TransactionUpdater, entityID, transactionID domain.ID, payload []byte) (Record[domain.Transaction], error) {
	if err := requireIDs(entityID, transactionID); err != nil {
		return Record[domain.Transaction]{}, err
	}
	if err := requireObject(payload); err != nil {
		return Record[domain.Transaction]{}, err
	}
	return repo.UpdateTransaction(ctx, entityID, transactionID, payload)
}
