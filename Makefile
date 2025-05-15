

DOCKER_COMPOSE="$(shell which docker-compose)"


restart: down up

up:
	${DOCKER_COMPOSE} up --build -d

down:
	${DOCKER_COMPOSE} down --remove-orphans


bash:
	${DOCKER_COMPOSE} exec postgres /bin/bash

contact:
	go run services/contact/cmd/app/main.go

m-up:
	go run migrations/auto.go

dump:
	docker exec -i clean-go-database psql -U user_db -d slurm < docker/dump/20220707173102_init_slurm.sql
	#docker run -d psql -U user_db -d clean-go_db < /home/ten/Videos/Cources/GO/slutm-capng/20220707173102_init_slurm.sql

