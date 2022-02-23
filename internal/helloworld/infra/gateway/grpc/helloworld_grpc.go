package grpc

import (
	pb "github.com/naokivandit/micro-service-docs/helloworld"
)

func Get() *pb.HelloReply {
	client := pb.NewGreeterClient(nil)

	req := &pb.HelloRequest{
		Name: "test_name",
	}

	res, err := client.SayHello(nil, req)
	if err != nil {
		return res
	}

	return res
}
