up:
	docker-compose -f docker-compose.yml up
	# docker-compose run --rm helloworld-api bash -c "go run cmd/main.go"

down:
	docker-compose -f docker-compose.yml down 