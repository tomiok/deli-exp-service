package engine

import (
	"github.com/deli/exp-service/datastore"
	"github.com/deli/exp-service/models"
)

type Spec interface {
	SaveWarehouse() string
	IndexDocument() string
	Search() []*models.ArticleResponse
	SearchById(uid string) *models.ArticleResponse
}

type Engine struct {
	datastore.ExperienceRepository
}

func (e *Engine) SaveWarehouse() string {
	e.ExperienceRepository.SaveWarehouse()
	return ""
}

func (e *Engine) IndexDocument() string {
	e.ExperienceRepository.IndexDocument()
	return ""
}
