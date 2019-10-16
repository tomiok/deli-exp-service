package datastore

import (
	"database/sql"
	"errors"
	"github.com/deli/exp-service/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
)

type SQLGateway interface {
	Save(exp models.ExperiencePost) (string, error)
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

func (sqlClient *SQLClient) Save(exp models.ExperiencePost) (string, error) {
	tx, err := sqlClient.Begin()

	if err != nil {
		return "", errors.New("cannot create transaction " + err.Error())
	}

	//stmt, err := tx.Prepare()
	//date

	_ = tx.Commit()

	return "", nil
}
