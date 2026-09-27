package application

import (
	"context"

	"github.com/dnutiu/fluxa-cli/internal/domain"
)

type AccountCreator interface {
	CreateAccount(context.Context, domain.ID, domain.AccountInput) (Record[domain.Account], error)
}

func AddAccount(ctx context.Context, repo AccountCreator, entityID domain.ID, input domain.AccountInput) (Record[domain.Account], error) {
	if err := requireIDs(entityID); err != nil {
		return Record[domain.Account]{}, err
	}
	if err := input.ValidateCreate(); err != nil {
		return Record[domain.Account]{}, err
	}
	return repo.CreateAccount(ctx, entityID, input)
}
