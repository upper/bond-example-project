DOCKER_CONTAINER := bond-db


DATABASE_HOST      ?= 127.0.0.1
DATABASE_NAME      ?= bond
DATABASE_PASSWORD  ?= postgres
DATABASE_PORT      ?= 5432
DATABASE_USER      ?= postgres

export DATABASE_HOST
export DATABASE_NAME
export DATABASE_PASSWORD
export DATABASE_PORT
export DATABASE_USER

run:
	cd service/web && go run *.go

test:
	make -C internal/tests test

docker-run:
	(docker rm -f $(DOCKER_CONTAINER) || exit 0) && \
	docker run \
		-d \
		--name $(DOCKER_CONTAINER) \
		-e POSTGRES_PASSWORD=$(DATABASE_PASSWORD) \
		-e POSTGRES_USER=$(DATABASE_USER) \
		-e POSTGRES_DB=$(DATABASE_NAME) \
		-p $(DATABASE_HOST):$(DATABASE_PORT):5432 \
		postgres && \
	sleep 5

db-load:
	docker exec -i $(DOCKER_CONTAINER) bash -c 'psql -U$(DATABASE_USER) $(DATABASE_NAME)' < sql/0000-init.sql

db-up: docker-run db-load
