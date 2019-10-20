package engine

import (
	"github.com/deli/exp-service/commons/logs"
	"github.com/deli/exp-service/datastore"
	"github.com/deli/exp-service/models"
)

type Spec interface {
	SaveWarehouse(exp models.ExperiencePost) (string, error)
	IndexDocument() error
	Search() []*models.ArticleResponse
	SearchById(uid string) *models.ArticleResponse
}

type Engine struct {
	document   datastore.ExperienceDocumentRepository
	relational datastore.ExperienceSQLRepository
}

func New(source string) Spec {
	sqlClient, err := datastore.NewMysqlClient(source)

	if err != nil {
		logs.Fatalf("cannot create mysql connection %s", err.Error())
		panic(err)
	}

	return &Engine{
		document:   &datastore.DocumentRepository{DC: &datastore.DocumentClient{}},
		relational: &datastore.SqlRepository{SqlGateway: sqlClient},
	}
}

func (e *Engine) SaveWarehouse(exp models.ExperiencePost) (string, error) {
	return e.relational.SaveWarehouse(exp)
}

func (e *Engine) IndexDocument() error {
	e.document.IndexDocument()
	return nil
}

func (e *Engine) Search() []*models.ArticleResponse {
	return nil
}

func (e *Engine) SearchById(uid string) *models.ArticleResponse {
	return nil
}
