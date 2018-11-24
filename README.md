# GIN SQLX CRUD

A tiny example to demonstrate implementing CRUD with gin and sqlx.

## Install dependencies

You should have go 1.11 and dep installed. Use dep to get dependencies;
```bash
dep ensure
```

## Run tests

I have just implemented minimal tests just to show how it can be done.
If you have go 1.11 installed use this to run unit tests;
```bash
go test github.com/wetterj/gin-sqlx-crud/...
```

If you have docker installed integration tests can be run with;
```bash
./sh/runIntegrationTests.sh
```

## Run a Development Env

To run the server locally use
```bash
go install github.com/wetterj/gin-sqlx-crud/cmd/srv

POSTGRES_HOST=localhost \
POSTGRES_PORT=22222 \
POSTGRES_USER=postgres \
POSTGRES_PASSWORD=testing_password \
POSTGRES_DB=postgres \
$GOPATH/bin/srv
```

Or using a script that creates a postgres DB in docker;
```bash
./sh/runDev.sh
```

## Deployment

To build a docker image of the server suitable for deployment into kubernetes use this;

```bash
docker build -t wetterj/gin-sqlx-crud -f dockerfiles/production.dockerfile .
```

There is a sample kubernetes manifest in `k8s/production.yaml`. It assumes
there is a docker registry available, a postgres DB available and a secret that describes the DB
user and password.

The service is exposed as ClusterIP. In my previous setups I used an nginx ingress network to
forward requests to the correct ClusterIP service, that way one AWS load balancer could
be used for multiple domains and services. Also nginx will be responsible for ssl
and compression etc.
