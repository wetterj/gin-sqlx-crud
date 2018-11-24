#!/bin/sh

# Launch a postgres DB for integration testing
docker run -d \
        --name gin-sqlx-crud-tests \
        -p 22222:5432 \
        -e POSTGRES_PASSWORD=testing_password \
        postgres

# Wait for DB to start up
# TODO: actually check readiness
sleep 2s

# Run the go tests
POSTGRES_HOST=localhost \
POSTGRES_PORT=22222 \
POSTGRES_USER=postgres \
POSTGRES_PASSWORD=testing_password \
POSTGRES_DB=postgres \
go test github.com/wetterj/gin-sqlx-crud/... -tags=integration -v

# Cleanup
docker rm -f gin-sqlx-crud-tests
