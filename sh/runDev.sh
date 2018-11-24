#!/bin/sh

# Launch a postgres DB for integration testing
docker run -d \
        --name gin-sqlx-crud-dev \
        -p 22222:5432 \
        -e POSTGRES_PASSWORD=testing_password \
        postgres

# Wait for DB to start up
# TODO: actually check readiness
sleep 2s

# compile
go install github.com/wetterj/gin-sqlx-crud/cmd/srv

# Run the go tests
POSTGRES_HOST=localhost \
POSTGRES_PORT=22222 \
POSTGRES_USER=postgres \
POSTGRES_PASSWORD=testing_password \
POSTGRES_DB=postgres \
$GOPATH/bin/srv

# Cleanup
docker rm -f gin-sqlx-crud-dev
