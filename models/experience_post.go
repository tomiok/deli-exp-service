package models

import (
	"strings"
	"time"
)

type ExperiencePost struct {
	UID       string  `json:"uid"`
	Title     string  `json:"title"`
	Subtitle  string  `json:"subtitle"`
	Body      string  `json:"body"`
	Tags      Tags    `json:"tags"`
	Product   Product `json:"product"`
	Published Publish `json:"published"`
}

type Publish struct {
	Date      time.Time `json:"date"`
	AuthorUID string    `json:"author"` //uid
}

type Product struct {
	UID     string    `json:"uid"`
	Name    string    `json:"name"`
	Date    time.Time `json:"date"`
	City    string    `json:"city"`
	Country string    `json:"country"`
	Details string    `json:"details"`
}

type Tags struct {
	UID  string   `json:"uid"`
	Tags []string `json:"tags"`
}

func (t *Tags) CsvValues() string {
	return strings.Join(t.Tags, ",")
}

// API response
type ArticleResponse struct {
	UID         string    `json:"uid"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Body        string    `json:"body"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Date        time.Time `json:"date"`
	AuthorName  string    `json:"author_name"`
	AuthorEmail string    `json:"author_email"`
}
