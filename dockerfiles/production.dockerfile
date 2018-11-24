FROM golang:1.11.1-alpine as builder

RUN apk update && apk upgrade && \
    apk add ca-certificates && \
    apk add --no-cache bash git openssh gcc musl-dev linux-headers

RUN go get -u -u github.com/golang/dep/cmd/dep

COPY . $GOPATH/src/github.com/wetterj/gin-sqlx-crud/
WORKDIR $GOPATH/src/github.com/wetterj/gin-sqlx-crud/

RUN $GOPATH/bin/dep ensure

# build the binary
RUN go install github.com/wetterj/gin-sqlx-crud/cmd/srv

# Copy to a smaller image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
# Copy our static executable
COPY --from=builder /go/bin/srv /go/bin/srv
ENTRYPOINT ["/go/bin/srv"]
