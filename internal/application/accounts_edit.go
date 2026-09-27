package application

import (
	"context"
	"fluxa-cli/internal/domain"
)

type AccountUpdater interface {
	UpdateAccount(context.Context, domain.ID, domain.ID, []byte) (Record[domain.Account], error)
}

func EditAccount(ctx context.Context, repo AccountUpdater, entityID, accountID domain.ID, payload []byte) (Record[domain.Account], error) {
	if err := requireIDs(entityID, accountID); err != nil {
		return Record[domain.Account]{}, err
	}
	if err := requireObject(payload); err != nil {
		return Record[domain.Account]{}, err
	}
	return repo.UpdateAccount(ctx, entityID, accountID, payload)
}
