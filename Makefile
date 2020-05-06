DOCKER_CONTAINER := bond-db


DB_HOST      ?= 127.0.0.1
DB_NAME      ?= bond
DB_PASSWORD  ?= postgres
DB_PORT      ?= 5432
DB_USER      ?= postgres

export DB_HOST
export DB_NAME
export DB_PASSWORD
export DB_PORT
export DB_USER

run:
	cd service/web && go run *.go

test:
	make -C internal/tests test

docker-run:
	(docker rm -f $(DOCKER_CONTAINER) || exit 0) && \
	docker run \
		-d \
		--name $(DOCKER_CONTAINER) \
		-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
		-e POSTGRES_USER=$(DB_USER) \
		-e POSTGRES_DB=$(DB_NAME) \
		-p $(DB_HOST):$(DB_PORT):5432 \
		postgres && \
	sleep 5

db-load:
	docker exec -i bond-db bash -c 'psql -Upostgres' < sql/0000-init.sql

db-up: docker-run db-load
