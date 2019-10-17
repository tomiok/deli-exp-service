package web

import (
	"github.com/deli/exp-service/commons"
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

func (er *ExperienceRequest) ToModel() models.ExperiencePost {

	tags := models.Tags{Tags: er.Tags}
	product := models.Product{
		Name:    er.Product.ProductName,
		Date:    time.Now(), //TODO parse this if the product is older. Now is not coming from the client
		City:    er.Product.City,
		Country: er.Product.Country,
		Details: er.Product.Details,
	}

	return models.ExperiencePost{
		UID:      commons.StringUUID(),
		Title:    er.Title,
		Subtitle: er.Subtitle,
		Body:     er.Body,
		Tags:     tags,
		Product:  product,
		Published: models.Publish{
			Date:      time.Now(),
			AuthorUID: er.AuthorUID,
		},
	}
}
