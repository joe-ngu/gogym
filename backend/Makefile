DOCKER_COMPOSE=docker compose
DOCKER_COMPOSE_FILE=docker-compose.yml
IMAGE_NAME=gogym

build:
	docker build --no-cache -t $(IMAGE_NAME) .

up:
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) up --build -d

down:
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) down

restart: down up

.PHONY: build up down restart
