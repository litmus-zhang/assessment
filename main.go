package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/litmus-zhang/assessment/api"
	"github.com/litmus-zhang/assessment/db"
)

const (
	dbDriver      = "postgres"
	dbUrl         = "postgresql://main:main@localhost:4000/main?sslmode=disable"
	ServerAddress = "localhost:8000"
)

func main() {
	var err error

	conn, err := sql.Open(dbDriver, dbUrl)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	db := db.New(conn)
	server := api.AppSetup(db)
	err = server.Start(ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
