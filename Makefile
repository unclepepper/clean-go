

DOCKER_COMPOSE="$(shell which docker-compose)"


restart: down up

up:
	${DOCKER_COMPOSE} up --build -d

down:
	${DOCKER_COMPOSE} down --remove-orphans

bash:
	${DOCKER_COMPOSE} exec postgres /bin/bash

ps:
	${DOCKER_COMPOSE} ps

logs:
	${DOCKER_COMPOSE} logs -f

contact:
	go run services/contact/cmd/app/main.go

m-up:
	go run services/contact/migrations/postgres/auto.go

dump:
	docker exec -i clean-go-database psql -U user_db -d postgres < docker/dump.sql

vendor:
	go mod vendor

tidy:
	go mod tidy


