package engine

import (
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
	datastore.ExperienceRepository
}

func (e *Engine) SaveWarehouse() (string, error) {
	e.ExperienceRepository.SaveWarehouse()
	return "", nil
}

func (e *Engine) IndexDocument() error {
	e.ExperienceRepository.IndexDocument()
	return nil
}
