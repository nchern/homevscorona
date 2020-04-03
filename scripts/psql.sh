#!/bin/sh
set -ue

DB_IMAGE="postgres:11-alpine"

DB_HOST="db"
DB_NAME="maindb"

DB_USER="root"

NETWORK="homevscorona_cor-net"

docker run --rm --network="$NETWORK" -it "$DB_IMAGE" psql -h "$DB_HOST" -U "$DB_USER" -d "$DB_NAME"
