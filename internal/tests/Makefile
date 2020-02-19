SHELL := bash

PROJECT      ?= example

DB_NAME      ?= example
DB_HOST      ?= 127.0.0.1
DB_PORT      ?= 5432
DB_USER      ?= postgres
DB_PASSWORD  ?= postgr3s

TEST_FLAGS ?=

export POSTGRES_VERSION

export DB_HOST
export DB_NAME
export DB_USER
export DB_PASSWORD
export DB_PORT

export TEST_FLAGS

server-up: server-down
	docker-compose -p $(PROJECT) up -d && \
	sleep 10

server-down:
	docker-compose -p $(PROJECT) down

require-client:
	@if [ -z "$$(which psql)" ]; then \
		echo 'Missing "psql" command. Please install the PostgreSQL client and try again.' && \
		exit 1; \
	fi

reset-db: require-client
	SQL="" && \
	SQL+="DROP DATABASE IF EXISTS $(DB_NAME);" && \
	SQL+="DROP ROLE IF EXISTS $(DB_USER);" && \
	SQL+="CREATE USER $(DB_USER) WITH PASSWORD '$(DB_PASSWORD)';" && \
	SQL+="CREATE DATABASE $(DB_NAME) ENCODING 'UTF-8' LC_COLLATE='en_US.UTF-8' LC_CTYPE='en_US.UTF-8' TEMPLATE template0;" && \
	SQL+="GRANT ALL PRIVILEGES ON DATABASE $(DB_NAME) TO $(DB_USER);" && \
	export PGPASSWORD="$(DB_PASSWORD)" && \
	psql -U$(DB_USER) -h$(DB_HOST) -p$(DB_PORT) template1 <<< $$SQL && \
	psql -U$(DB_USER) -h$(DB_HOST) -p$(DB_PORT) $(DB_NAME) < ../../model/schema/*.sql;

test: reset-db
	go test $(TEST_FLAGS) -v ./...

