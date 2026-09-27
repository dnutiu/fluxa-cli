package application

import (
	"context"
	"fluxa-cli/internal/domain"
)

type SubscriptionActivator interface {
	SetSubscriptionActive(context.Context, domain.ID, domain.ID, bool) (Record[domain.Subscription], error)
}

func SetSubscriptionActive(ctx context.Context, repo SubscriptionActivator, entityID, subscriptionID domain.ID, active bool) (Record[domain.Subscription], error) {
	if err := requireIDs(entityID, subscriptionID); err != nil {
		return Record[domain.Subscription]{}, err
	}
	return repo.SetSubscriptionActive(ctx, entityID, subscriptionID, active)
}
