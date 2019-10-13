package engine

import "github.com/deli/exp-service/models"

type Spec interface {
	SaveWarehouse() string
	IndexDocument() string
	Search() []*models.ArticleResponse
	SearchById(uid string) *models.ArticleResponse
}

type Engine struct {
}

func (e *Engine) SaveWarehouse() string {
	return ""
}

func (e *Engine) IndexDocument() string {

	return ""
}
