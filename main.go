package main

import (
	"database/sql"
	"log"

	"github.com/isaac8838/simplebank/api"
	db "github.com/isaac8838/simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:password@host.docker.internal:5432/simple_bank?sslmode=disable"
	serverAddress = "localhost:8080"
)

func main() {
	var err error

	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
