FROM golang:1.22.3-alpine

WORKDIR /app/reliab-test

RUN apk add --no-cache postgresql-client build-base
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

