.PHONY: build up down restart logs clean ps exec migrate env-setup help

# Docker compose command with options
DOCKER_COMPOSE := docker-compose

# Colors
GREEN := \033[0;32m
YELLOW := \033[0;33m
NC := \033[0m # No Color

help:
	@echo "${YELLOW}ScoreMaster Docker Management${NC}"
	@echo "${GREEN}make build${NC}    - Build all containers"
	@echo "${GREEN}make up${NC}       - Start all containers in detached mode"
	@echo "${GREEN}make down${NC}     - Stop and remove all containers"
	@echo "${GREEN}make restart${NC}  - Restart all containers"
	@echo "${GREEN}make logs${NC}     - View logs of all containers"
	@echo "${GREEN}make logs-api${NC} - View logs of the API container"
	@echo "${GREEN}make logs-db${NC}  - View logs of the database container"
	@echo "${GREEN}make clean${NC}    - Stop containers and remove volumes"
	@echo "${GREEN}make ps${NC}       - List running containers"
	@echo "${GREEN}make exec-api${NC} - Execute commands in API container"
	@echo "${GREEN}make exec-db${NC}  - Execute commands in DB container"

build:
	@echo "${YELLOW}Building containers...${NC}"
	@$(DOCKER_COMPOSE) build

up:
	@echo "${YELLOW}Starting containers in detached mode...${NC}"
	@$(DOCKER_COMPOSE) up -d

down:
	@echo "${YELLOW}Stopping containers...${NC}"
	@$(DOCKER_COMPOSE) down

restart:
	@echo "${YELLOW}Restarting containers...${NC}"
	@$(DOCKER_COMPOSE) restart

logs:
	@echo "${YELLOW}Viewing logs...${NC}"
	@$(DOCKER_COMPOSE) logs -f

logs-api:
	@echo "${YELLOW}Viewing API logs...${NC}"
	@$(DOCKER_COMPOSE) logs -f api

logs-db:
	@echo "${YELLOW}Viewing database logs...${NC}"
	@$(DOCKER_COMPOSE) logs -f db

clean:
	@echo "${YELLOW}Stopping containers and removing volumes...${NC}"
	@$(DOCKER_COMPOSE) down -v

ps:
	@echo "${YELLOW}Listing running containers...${NC}"
	@$(DOCKER_COMPOSE) ps

exec-api:
	@echo "${YELLOW}Executing command in API container...${NC}"
	@$(DOCKER_COMPOSE) exec api sh

exec-db:
	@echo "${YELLOW}Executing command in DB container...${NC}"
	@$(DOCKER_COMPOSE) exec db bash