# Variables
DOCKER_COMPOSE=docker compose
DOCKER_COMPOSE_FILE=docker-compose.yml
IMAGE_NAME=gogym

# Build the Docker image
build:
	docker build -t $(IMAGE_NAME) .

# Bring up the Docker Compose services
up:
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) up --build

# Tear down the Docker Compose services
down:
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) down

# Restart the Docker Compose services
restart: down up

# Phony targets to avoid conflicts with files of the same name
.PHONY: build up down restart
