package application

import (
	"context"
	"fluxa-cli/internal/domain"
)

type TransactionDeleter interface {
	DeleteTransaction(context.Context, domain.ID, domain.ID) error
}

func DeleteTransaction(ctx context.Context, repo TransactionDeleter, entityID, transactionID domain.ID) error {
	if err := requireIDs(entityID, transactionID); err != nil {
		return err
	}
	return repo.DeleteTransaction(ctx, entityID, transactionID)
}
