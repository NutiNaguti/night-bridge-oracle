package db

import (
	"log"

	"github.com/jackc/pgx"
	"github.com/shopspring/decimal"
)

type Transactions struct {
	id        int
	from      string
	to        string
	amount    decimal.Decimal
	timestamp int64
}

func SetupConnection(host string, port uint16, database string, user string, password string) *pgx.Conn {
	connConfig := &pgx.ConnConfig{
		Host:     host,
		Port:     port,
		Database: database,
		User:     user,
		Password: password,
	}
	conn, err := pgx.Connect(*connConfig)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func NewTransaction() {
}
