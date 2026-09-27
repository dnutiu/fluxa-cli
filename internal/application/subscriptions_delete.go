package application

import (
	"context"
	"fluxa-cli/internal/domain"
)

type SubscriptionDeleter interface {
	DeleteSubscription(context.Context, domain.ID, domain.ID) error
}

func DeleteSubscription(ctx context.Context, repo SubscriptionDeleter, entityID, subscriptionID domain.ID) error {
	if err := requireIDs(entityID, subscriptionID); err != nil {
		return err
	}
	return repo.DeleteSubscription(ctx, entityID, subscriptionID)
}
