package datastore

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
)

type SQLGateway interface {
	Save()
}

type SQLClient struct {
	*sql.DB
}

func NewMysqlDS(source string) (*SQLClient, error) {

	connection, err := sql.Open("mysql", source)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &SQLClient{connection}, nil
}

func (sqlClient *SQLClient) Save() {
	tx, err := sqlClient.Begin()

	if err != nil {

	}
	_ = tx.Commit()
}
