#!/bin/sh
set -ue

FLYWAY_IMAGE="flyway/flyway:6.0.7-alpine"

MIGRATIONS_PATH="backend/maindb/migrations/"

DB_HOST="db"
DB_NAME="maindb"

DB_USER="root"
DB_PASSWORD=${DB_PASSWORD:-"root"}


NETWORK="homevscorona_cor-net"

DB_URL="jdbc:postgresql://$DB_HOST/$DB_NAME"


docker run --rm \
    --network="$NETWORK" \
    -v "$(realpath $MIGRATIONS_PATH):/flyway/sql" \
    -it "$FLYWAY_IMAGE" \
    -url="$DB_URL" \
    -user="$DB_USER" \
    -password="$DB_PASSWORD" \
    migrate
