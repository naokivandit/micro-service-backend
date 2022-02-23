package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/naokivandit/gqlgen-todos/graph/model"
)

func (r *mutationResolver) CreateOrder(ctx context.Context, dryRun bool) (*model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteOrder(ctx context.Context, dryRun bool) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Order(ctx context.Context) (*model.Order, error) {
	order := &model.Order{}
	return order, nil
}

func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}
