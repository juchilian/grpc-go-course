DOCKER_YAML=-f docker-compose.yml

DOCKER=COMPOSE_PROJECT_NAME=grpc-project docker-compose $(DOCKER_YAML)

generate:
	$(DOCKER) run --rm go-test ./scripts/generate.sh

generate2:
	$(DOCKER) run --rm go-test protoc greet/greetpb/greet.proto --go_out=plugins=grpc:. 