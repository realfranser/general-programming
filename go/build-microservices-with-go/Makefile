# Constants
SHELL := /bin/bash
## Postgres
POSTGRES_USER := user
POSTGRES_PASSWORD := password
POSTGRES_PORT := 5432

install-pipelines:
	go install github.com/MarioCarrion/complex-pipelines/part5

install-migrate:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.14.1

install-tools: install-pipelines intall-migrate

migrate-tables-up:
	migrate -path ./db/migrations/ -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/dbname?sslmode=disable" up

migrate-tables-down:
	migrate -path ./db/migrations/ -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/dbname?sslmode=disable" down

launch-app:
	@bash ./scripts/launch.sh

run:
	go run .

direnv-allow:
	direnv allow .

docker-kill:
	docker kill memcache postgres_db
