package presentaion

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	generated1 "github.com/naokivandit/micro-service-backend/helloworld_api/api/graphql/generated"
	model1 "github.com/naokivandit/micro-service-backend/helloworld_api/internal/domain/model"
)

func (r *queryResolver) SayHello(ctx context.Context, name *string) (*model1.HelloReply, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
