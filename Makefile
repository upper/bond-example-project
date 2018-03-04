SHELL := bash

DB_HOST   ?= 127.0.0.1
DB_PORT   ?= 5432

DB_USERNAME ?= upperio_tests
DB_PASSWORD ?= upperio_secret
DB_NAME     ?= upperio_tests

TEST_FLAGS ?=

export DB_HOST
export DB_NAME
export DB_PASSWORD
export DB_PORT
export DB_USERNAME

build:
	go build && go install

require-client:
	@if [ -z "$$(which psql)" ]; then \
		echo 'Missing "psql" command. Please install the PostgreSQL client and try again.' && \
		exit 1; \
	fi

reset-db: require-client
	SQL="" && \
	SQL+="DROP DATABASE IF EXISTS $(DB_NAME);" && \
	SQL+="DROP ROLE IF EXISTS $(DB_USERNAME);" && \
	SQL+="CREATE USER $(DB_USERNAME) WITH PASSWORD '$(DB_PASSWORD)';" && \
	SQL+="CREATE DATABASE $(DB_NAME) ENCODING 'UTF-8' LC_COLLATE='en_US.UTF-8' LC_CTYPE='en_US.UTF-8' TEMPLATE template0;" && \
	SQL+="GRANT ALL PRIVILEGES ON DATABASE $(DB_NAME) TO $(DB_USERNAME);" && \
	if [ ! -z "$$CI" ]; then \
		psql -Upostgres -h$(DB_HOST) -p$(DB_PORT) <<< $$SQL && \
		psql -Upostgres -h$(DB_HOST) -p$(DB_PORT) < model/schema/*.sql; \
	else \
		export PGPASSWORD="$(DB_PASSWORD)" && \
		psql -U$(DB_USERNAME) -h$(DB_HOST) -p$(DB_PORT) <<< $$SQL && \
		psql -U$(DB_USERNAME) -h$(DB_HOST) -p$(DB_PORT) < model/schema/*.sql; \
	fi

test: reset-db
	go test $(TEST_FLAGS) -v ./...
