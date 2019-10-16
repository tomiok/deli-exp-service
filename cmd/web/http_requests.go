package web

import (
	"github.com/deli/exp-service/models"
	"time"
)

//TODO add all the DTOs to map with json request here
type ExperienceRequest struct {
	Title     string         `json:"title"`
	Subtitle  string         `json:"subtitle"`
	Body      string         `json:"body"`
	AuthorUID string         `json:"author_uid"`
	Product   ProductRequest `json:"product"`
	Tags      []string       `json:"tags"`
}

type ProductRequest struct {
	ProductName string    `json:"product_name"`
	Details     string    `json:"details"`
	Date        time.Time `json:"date"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
}

func (er *ExperienceRequest) ToModel() *models.ExperiencePost {
	return &models.ExperiencePost{
		UID:       "",
		Title:     "",
		Subtitle:  "",
		Body:      "",
		Tags:      models.Tags{},
		Product:   models.Product{},
		Published: models.Publish{},
	}
}
