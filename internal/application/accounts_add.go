package application

import (
	"context"
	"fluxa-cli/internal/domain"
)

type AccountCreator interface {
	CreateAccount(context.Context, domain.ID, []byte) (Record[domain.Account], error)
}

func AddAccount(ctx context.Context, repo AccountCreator, entityID domain.ID, payload []byte) (Record[domain.Account], error) {
	if err := requireIDs(entityID); err != nil {
		return Record[domain.Account]{}, err
	}
	if err := requireObject(payload); err != nil {
		return Record[domain.Account]{}, err
	}
	return repo.CreateAccount(ctx, entityID, payload)
}
