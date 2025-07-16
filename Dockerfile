FROM golang:1.24.5-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY *.go ./
COPY static/ static/
COPY views/ views/

RUN go mod tidy
RUN go build -o app .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/app .
COPY config.json .

ENTRYPOINT ["./app"]
