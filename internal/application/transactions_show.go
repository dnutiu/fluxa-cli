package application

import (
	"context"
	"fluxa-cli/internal/domain"
)

type TransactionGetter interface {
	GetTransaction(context.Context, domain.ID, domain.ID) (Record[domain.Transaction], error)
}

func ShowTransaction(ctx context.Context, repo TransactionGetter, entityID, transactionID domain.ID) (Record[domain.Transaction], error) {
	if err := requireIDs(entityID, transactionID); err != nil {
		return Record[domain.Transaction]{}, err
	}
	return repo.GetTransaction(ctx, entityID, transactionID)
}
