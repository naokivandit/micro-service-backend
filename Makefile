gqlgen:
	go run github.com/99designs/gqlgen generate



server:
	go run cmd/helloworld_client/main.go

helloworldserver:
	go run cmd/helloworld_server/main.go
