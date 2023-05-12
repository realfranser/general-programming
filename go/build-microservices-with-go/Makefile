# Constants
SHELL := /bin/bash
## Postgres
POSTGRES_USER := user
POSTGRES_PASSWORD := password
POSTGRES_DB := dbname
POSTGRES_PORT := 5432
# Memcache
MEMCACHED_PORT := 11211

postgres-docker-run:
	docker run \
    -d \
    -e POSTGRES_HOST_AUTH_METHOD=trust \
    -e POSTGRES_USER=$(POSTGRES_USER) \
    -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
    -e POSTGRES_DB=$(POSTGRES_DB) \
    -p $(POSTGRES_PORT):$(POSTGRES_PORT) \
	  postgres:12.5-alpine

memcached-docker-run:
	docker run \
  	-d \
    -p $(MEMCACHED_PORT):$(MEMCACHED_PORT) \
    memcached:1.6.9-alpine

install-pipelines:
	go install github.com/MarioCarrion/complex-pipelines/part5

install-migrate:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.14.1

migrate-tables-up:
	migrate -path ./db/migrations/ -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/dbname?sslmode=disable" up

migrate-tables-down:
	migrate -path ./db/migrations/ -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/dbname?sslmode=disable" down

run-data-pipelines:
	DATABASE_URL="postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/dbname?sslmode=disable" part5

launch-app:
	@bash ./scripts/launch.sh

run:
	go run .

direnv-allow:
	direnv allow .

docker-kill:
	docker kill memcache postgres_db
