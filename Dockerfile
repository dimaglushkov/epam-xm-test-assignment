FROM golang:1.20.3-alpine3.17 as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

FROM golang:1.20.3-alpine3.17 as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN go build -o /bin/app ./cmd/

FROM scratch
COPY --from=builder /app/migrations /migrations
COPY --from=builder /bin/app /app
CMD ["/app"]