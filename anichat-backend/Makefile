# Backend development helper commands

COMPOSE_FILE=docker-compose.yml

.PHONY: up down rebuild logs ps shell migrate

up:
	docker compose -f $(COMPOSE_FILE) up -d --build

down:
	docker compose -f $(COMPOSE_FILE) down -v
	
restart: down up

migrate:
	docker compose -f $(COMPOSE_FILE) run --rm migrate
