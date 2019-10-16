package datastore

import (
	"errors"
	"github.com/deli/exp-service/models"
)

type ExperienceSQLRepository interface {
	SaveWarehouse(exp models.ExperiencePost) (string, error)
}

type ExperienceDocumentRepository interface {
	IndexDocument()
}

type ExperienceRepository interface {
	ExperienceDocumentRepository
	ExperienceSQLRepository
}

type SqlRepository struct {
	SQLGateway
}

type DocumentRepository struct {
	DC DocIndexClient
}

func (sqlRepo *SqlRepository) SaveWarehouse(exp models.ExperiencePost) (string, error) {

	//TODO map here model based struct to SQL based struct
	uid, err := sqlRepo.Save(exp)

	if err != nil {
		return "", errors.New("cannot save experience, " + err.Error())
	}

	return uid, nil
}

func (docClient *DocumentRepository) IndexDocument() {
	docClient.DC.Index()
}
