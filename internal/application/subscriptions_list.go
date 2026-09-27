package application

import (
	"context"
	"github.com/dnutiu/fluxa-cli/internal/domain"
)

type SubscriptionLister interface {
	ListSubscriptions(context.Context, domain.ID, domain.SubscriptionFilter) (Collection[domain.Subscription], error)
}

func ListSubscriptions(ctx context.Context, repo SubscriptionLister, entityID domain.ID, filter domain.SubscriptionFilter) (Collection[domain.Subscription], error) {
	if err := requireIDs(entityID); err != nil {
		return Collection[domain.Subscription]{}, err
	}
	if err := filter.Validate(); err != nil {
		return Collection[domain.Subscription]{}, err
	}
	return repo.ListSubscriptions(ctx, entityID, filter)
}
