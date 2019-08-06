FROM golang:1.12 as builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
WORKDIR /go/src/github.com/kazukousen/go-api-utils/
COPY . .
RUN SHORT_SHA=$(git rev-parse --short HEAD) && \
    go build -o bin/app -a -tags netgo -installsuffix netgo -ldflags="-s -w -X \"main.ShortSHA=${SHORT_SHA}\" -extldflags -static" main.go

FROM alpine:latest
EXPOSE 8080
EXPOSE 9200
# to use connection with SSL/TLS
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/kazukousen/go-api-utils/bin/app /app
CMD ["/app"]
