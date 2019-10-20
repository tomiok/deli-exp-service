package datastore

import (
	"database/sql"
	"errors"
	"github.com/deli/exp-service/commons/logs"
	"github.com/deli/exp-service/datastore/storage"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
)

type SQLGateway interface {
	Save(exp *storage.ExperienceDTO) (string, error)
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

func (sqlClient *SQLClient) Save(exp *storage.ExperienceDTO) (string, error) {
	tx, err := sqlClient.Begin()

	if err != nil {
		return "", errors.New("cannot create transaction " + err.Error())
	}

	tags := exp.SqlTags
	product := exp.SqlProduct
	experience := exp.SQLExperiencePost
	err = inTransactionTags(tx, &tags)
	err = inTransactionProduct(tx, &product)
	err = inTransactionExperience(tx, &experience)

	if err != nil {
		logs.Errorf("errors in transaction %s", err.Error())
		err = tx.Rollback()
		return "", err
	}

	_ = tx.Commit()

	return experience.UID, nil
}

func inTransactionExperience(tx *sql.Tx, exp *storage.SQLExperiencePost) error {
	stmt, err := tx.Prepare(`insert into experience_post (uid, title, subtitle, body, date, author_uid, tag_uid,
                             product_uid, photo_url) values (?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		logs.Errorf("cannot prepare statement %s", err.Error())
		return err
	}

	defer func() {
		_ = stmt.Close()
	}()

	res, err := stmt.Exec(exp.UID,
		exp.Title,
		exp.Subtitle,
		exp.Body,
		exp.Date,
		exp.AuthorUID,
		exp.TagUID,
		exp.ProductUID,
		exp.PhotoURL)

	if err != nil {
		logs.Errorf("cannot execute statement %s", err.Error())
		return err
	}
	rows, err := res.RowsAffected()
	logs.Infof("rows affected: %d", rows)

	return err
}

func inTransactionProduct(tx *sql.Tx, product *storage.SQLProduct) error {
	stmt, err := tx.Prepare("insert into product (uid, name, date, city, country, details) values(?,?,?,?,?,?)")

	if err != nil {
		logs.Errorf("cannot prepare SQL statement %s", err.Error())
		return err
	}

	defer func() {
		_ = stmt.Close()
	}()

	res, err := stmt.Exec(product.UID, product.Name, product.Date, product.City, product.Country, product.Details)

	if err != nil {
		logs.Errorf("cannot execute statement %s", err.Error())
		return err
	}
	rows, err := res.RowsAffected()
	logs.Infof("rows affected: %d", rows)

	return err
}

func inTransactionTags(tx *sql.Tx, tags *storage.SQLTags) error {
	stmt, err := tx.Prepare("insert into tags (uid, csv_tags) values(?,?)")

	if err != nil {
		return err
	}

	defer func() {
		_ = stmt.Close()
	}()

	res, err := stmt.Exec(tags.UID, tags.Tags)

	if err != nil {
		logs.Errorf("cannot execute statement %s", err.Error())
		return err
	}

	rows, err := res.RowsAffected()

	logs.Infof("rows affected: %d", rows)

	return err
}
