package datastore

import (
	"errors"
	"github.com/deli/exp-service/datastore/storage"
	"github.com/deli/exp-service/models"
)

type ExperienceSQLRepository interface {
	SaveWarehouse(exp models.ExperiencePost) (string, error)
}

type ExperienceDocumentRepository interface {
	IndexDocument()
}

type SqlRepository struct {
	SqlGateway SQLGateway
}

type DocumentRepository struct {
	DC DocIndexClient
}

func (sqlRepo *SqlRepository) SaveWarehouse(exp models.ExperiencePost) (string, error) {
	expSQL := storage.FromProduct(exp)

	uid, err := sqlRepo.SqlGateway.Save(expSQL)

	if err != nil {
		return "", errors.New("cannot save experience, " + err.Error())
	}

	return uid, nil
}

func (docClient *DocumentRepository) IndexDocument() {
	docClient.DC.Index()
}
