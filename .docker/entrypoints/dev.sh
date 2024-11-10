#!/bin/sh

until PGHOST=${DB_HOST} PGDATABASE=${DB_NAME} PGPORT=${DB_PORT} PGUSER=${DB_USER} PGPASSWORD=${DB_PASS} psql -c 'SELECT 1'; do sleep 5; done;

echo "*** Begin migrate ***"

echo "appenv=$APP_ENV"

export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable
export GOOSE_MIGRATION_DIR=./cmd/migrations
goose -allow-missing up

echo "*** End migrate ***"
echo "Removing main if exists..."
if [ -f /tmp/main ]; then
    rm /tmp/main
fi

echo "Building main..."
go build -mod=vendor -race -o /tmp/main ./cmd/main.go
echo "Starting main..."
exec /tmp/main