package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"time"
)

type connectorManager interface {
	getConnection() (*sqlx.DB, error)
	closeConnection(conn *sqlx.DB)
}

var _ connectorManager = (*DatabaseConnectionManager)(nil)

type DatabaseConnectionManager struct{}

func (r DatabaseConnectionManager) getConnection() (*sqlx.DB, error) {
	uri, err := getPostgresConnectionURI()
	if err != nil {
		return nil, err
	}

	db, err := sqlx.Open("postgres", uri)

	if err != nil {
		log.Print("Error while establishing db connection: " + err.Error())
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}

func (r DatabaseConnectionManager) closeConnection(conn *sqlx.DB) {
	err := conn.Close()
	if err != nil {
		log.Error().Err(err)
	}
}
