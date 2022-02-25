package presentaion

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/naokivandit/micro-service-backend/helloworld_api/api/graphql/generated"
	"github.com/naokivandit/micro-service-backend/helloworld_api/internal/domain/model"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

func (r *queryResolver) SayHello(ctx context.Context, name *string) (*model.HelloReply, error) {
	flag.Parse()
	// Set up a connection to the server.
	// conn, err := grpc.Dial(*addr, grpc.WithInsecure())

	conn, err := grpc.Dial("helloworld:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()
	res, err := c.SayHello(ctx, &pb.HelloRequest{Name: "name"})
	if err != nil {
		return &model.HelloReply{
			Message: "error",
		}, nil
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", res.GetMessage())

	return &model.HelloReply{
		Message: res.Message,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
