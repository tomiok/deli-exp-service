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

func (e *Engine) SaveWarehouse(exp models.ExperiencePost) (string, error) {
	return e.ExperienceRepository.SaveWarehouse(exp)
}

func (e *Engine) IndexDocument() error {
	e.ExperienceRepository.IndexDocument()
	return nil
}

func (e *Engine) Search() []*models.ArticleResponse {
	return nil
}

func (e *Engine) SearchById(uid string) *models.ArticleResponse {
	return nil
}
