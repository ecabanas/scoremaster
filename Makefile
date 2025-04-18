DOCKER_COMPOSE := docker-compose

# Colors
GREEN := \033[0;32m
YELLOW := \033[0;33m
NC := \033[0m # No Color

help:
	@echo "${YELLOW}ScoreMaster Docker Management${NC}"
	@echo "${GREEN}make build${NC}    		- Build all containers"
	@echo "${GREEN}make up${NC}       		- Start all containers in detached mode"
	@echo "${GREEN}make down${NC}     		- Stop and remove all containers"
	@echo "${GREEN}make re${NC}  		- Restart all containers"
	@echo "${GREEN}make logs${NC}     		- View logs of all containers"
	@echo "${GREEN}make logs-front${NC}     	- View logs of the frontend container"
	@echo "${GREEN}make logs-back${NC}		- View logs of the backend container"
	@echo "${GREEN}make logs-db${NC}  		- View logs of the database container"
	@echo "${GREEN}make clean${NC}    		- Stop containers and remove volumes"
	@echo "${GREEN}make ps${NC}       		- List running containers"
	@echo "${GREEN}make exec-front${NC} 	- Execute commands in frontend container"
	@echo "${GREEN}make exec-back${NC} 	- Execute commands in backend container"
	@echo "${GREEN}make exec-db${NC}  		- Execute commands in DB container"

build:
	@echo "${YELLOW}Building containers...${NC}"
	@$(DOCKER_COMPOSE) build

up:
	@echo "${YELLOW}Starting containers in detached mode...${NC}"
	@$(DOCKER_COMPOSE) up -d

down:
	@echo "${YELLOW}Stopping containers...${NC}"
	@$(DOCKER_COMPOSE) down

re:
	@echo "${YELLOW}Restarting containers...${NC}"
	@$(DOCKER_COMPOSE) restart

logs:
	@echo "${YELLOW}Viewing logs...${NC}"
	@$(DOCKER_COMPOSE) logs -f

logs-front:
	@echo "${YELLOW}Viewing frontend logs...${NC}"
	@$(DOCKER_COMPOSE) logs -f frontend

logs-back:
	@echo "${YELLOW}Viewing backend logs...${NC}"
	@$(DOCKER_COMPOSE) logs -f backend

logs-db:
	@echo "${YELLOW}Viewing database logs...${NC}"
	@$(DOCKER_COMPOSE) logs -f db

clean:
	@echo "${YELLOW}Stopping containers and removing volumes...${NC}"
	@$(DOCKER_COMPOSE) down -v

ps:
	@echo "${YELLOW}Listing running containers...${NC}"
	@$(DOCKER_COMPOSE) ps

exec-front:
	@echo "${YELLOW}Executing command in frontend container...${NC}"
	@$(DOCKER_COMPOSE) exec frontend sh

exec-back:
	@echo "${YELLOW}Executing command in backend container...${NC}"
	@$(DOCKER_COMPOSE) exec backend sh

exec-db:
	@echo "${YELLOW}Executing command in DB container...${NC}"
	@$(DOCKER_COMPOSE) exec db bash

.PHONY: build up down restart logs clean ps exec migrate env-setup help
