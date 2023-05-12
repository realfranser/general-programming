#!/bin/bash

# Constants
## Postgres
readonly POSTGRES_HOST_AUTH_METHOD="trust"
readonly POSTGRES_USER="user"
readonly POSTGRES_PASSWORD="password"
readonly POSTGRES_DB="dbname"
readonly POSTGRES_PORT=5432
readonly POSTGRES_IMAGE="postgres:12.5-alpine"

## DB
readonly DB_MIGRATIONS_PATH="db/migrations/"
readonly DB_CONTAINER_NAME="postgres_db"

## Memcache
readonly MEMCACHE_PORT=11211
readonly MEMCACHE_IMAGE="memcached:1.6.9-alpine"
readonly MEMCACHE_CONTAINER_NAME="memcache"

function launch_postgres_container() {
    if [ $( docker ps -a | grep "\<$DB_CONTAINER_NAME\>" | wc -l ) -gt 0 ]; then
        docker start $DB_CONTAINER_NAME
    else
        docker run \
          -d \
          --name $DB_CONTAINER_NAME \
          -e POSTGRES_HOST_AUTH_METHOD=$POSTGRES_HOST_AUTH_METHOD \
          -e POSTGRES_USER=$POSTGRES_USER \
          -e POSTGRES_PASSWORD=$POSTGRES_PASSWORD \
          -e POSTGRES_DB=$POSTGRES_DB \
          -p $POSTGRES_PORT:$POSTGRES_PORT \
          $POSTGRES_IMAGE
    fi
}

function launch_memcache_container() {
    if [ $( docker ps -a | grep "\<$MEMCACHE_CONTAINER_NAME\>" | wc -l ) -gt 0 ]; then
        docker start $MEMCACHE_CONTAINER_NAME
    else
        docker run \
          -d \
          --name $MEMCACHE_CONTAINER_NAME \
          -p $MEMCACHE_PORT:$MEMCACHE_PORT\
          $MEMCACHE_IMAGE
    fi
}

function database_migrations() {
    ./scripts/custom_wait.sh -t 10 localhost:$POSTGRES_PORT -- echo "DB is up"
    tries=10

    while [ "$tries" -gt 0 ]; do
        if [[ $(migrate -path $DB_MIGRATIONS_PATH -database $DATABASE_URL up) == 'no change' ]]
        then
            echo "Tables migrated correctly"
            break
        else
            sleep 1
        fi

        tries=$(( tries - 1 ))
    done

    if [ "$tries" -eq 0 ]; then
        echo 'failed when migrating db table schemas' >&2
        exit 1
    fi
}

function populate_db() {
    part5
}

launch_postgres_container
launch_memcache_container
database_migrations
# populate_db
