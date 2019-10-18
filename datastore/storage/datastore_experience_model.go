package storage

import (
	"github.com/deli/exp-service/models"
	"time"
)

type SQLExperienceDTO struct {
	SQLExperiencePost
	SQLTags
	SQLProduct
}

type SQLExperiencePost struct {
	UID        string
	Title      string
	Subtitle   string
	Body       string
	Date       time.Time
	TagsUID    string
	AuthorUID  string
	ProductUID string
	PhotoURL   string
}

type SQLTags struct {
	UID  string
	tags string
}

type SQLProduct struct {
	UID     string
	Name    string
	Details string
	Date    time.Time
	City    string
	Country string
}

func FromProduct(post models.ExperiencePost) *SQLExperienceDTO {
	return &SQLExperienceDTO{
		SQLExperiencePost: SQLExperiencePost{
			UID:        post.UID,
			Title:      post.Title,
			Subtitle:   post.Subtitle,
			Body:       post.Body,
			Date:       post.Published.Date,
			TagsUID:    post.Tags.UID,
			AuthorUID:  post.Published.AuthorUID,
			ProductUID: post.Product.UID,
			PhotoURL:   "", //TODO finish this
		},
		SQLTags: SQLTags{
			UID:  post.Tags.UID,
			tags: post.Tags.CsvValues(),
		},
		SQLProduct: SQLProduct{
			UID:     post.Product.UID,
			Name:    post.Product.Name,
			Details: post.Product.Details,
			Date:    post.Product.Date,
			City:    post.Product.City,
			Country: post.Product.Country,
		},
	}
}
