package application

import (
	"context"
	"errors"
	"github.com/dnutiu/fluxa-cli/internal/domain"
	"strings"
)

type TransactionCreator interface {
	CreateTransaction(context.Context, domain.ID, []byte, string) (Record[domain.Transaction], error)
}

func AddTransaction(ctx context.Context, repo TransactionCreator, entityID domain.ID, payload []byte, key string) (Record[domain.Transaction], error) {
	if err := requireIDs(entityID); err != nil {
		return Record[domain.Transaction]{}, err
	}
	if err := requireObject(payload); err != nil {
		return Record[domain.Transaction]{}, err
	}
	if strings.TrimSpace(key) == "" || len(key) > 200 {
		return Record[domain.Transaction]{}, errors.New("idempotency key must be 1 to 200 characters")
	}
	return repo.CreateTransaction(ctx, entityID, payload, key)
}
