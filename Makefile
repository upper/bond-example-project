DB_HOST      ?= 127.0.0.1
DB_NAME      ?= example
DB_PASSWORD  ?= postgr3s
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
