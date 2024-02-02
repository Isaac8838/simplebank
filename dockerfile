FROM golang:1.21.6-bookworm

RUN apt-get update && \
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /simple_bank

COPY . .

EXPOSE 8888
