package engine

import "deli/exp-service/model"

type Spec interface {
	SaveWarehouse() string
	IndexDocument() string
	Search() [] *model.ArticleResponse
	SearchById(uid string) *model.ArticleResponse
}

type Engine struct {
}

func (e *Engine) SaveWarehouse() string {
	return ""
}

func (e *Engine) IndexDocument() string {

	return ""
}
