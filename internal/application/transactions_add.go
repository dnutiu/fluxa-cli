package application

import (
	"context"
	"errors"
	"strings"

	"github.com/dnutiu/fluxa-cli/internal/domain"
)

type TransactionCreator interface {
	CreateTransaction(context.Context, domain.ID, domain.TransactionInput, string) (Record[domain.Transaction], error)
}

func AddTransaction(ctx context.Context, repo TransactionCreator, entityID domain.ID, input domain.TransactionInput, key string) (Record[domain.Transaction], error) {
	if err := requireIDs(entityID); err != nil {
		return Record[domain.Transaction]{}, err
	}
	if err := input.ValidateCreate(); err != nil {
		return Record[domain.Transaction]{}, err
	}
	if strings.TrimSpace(key) == "" || len(key) > 200 {
		return Record[domain.Transaction]{}, errors.New("idempotency key must be 1 to 200 characters")
	}
	return repo.CreateTransaction(ctx, entityID, input, key)
}
