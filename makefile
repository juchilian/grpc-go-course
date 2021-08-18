DOCKER_YAML=-f docker-compose.yml

DOCKER=COMPOSE_PROJECT_NAME=grpc-go-course docker-compose $(DOCKER_YAML)

exec-redis:
	$(DOCKER) exec redis sh
	# docker-compose exec redis sh

generate:
	$(DOCKER) run --rm go-test ./scripts/generate.sh

generate2:
	$(DOCKER) run --rm go-test protoc greet/greetpb/greet.proto --go_out=plugins=grpc:. 