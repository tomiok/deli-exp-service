package models

import "time"

type ExperiencePost struct {
	Uid       string   `json:"uid"`
	Title     string   `json:"title"`
	Subtitle  string   `json:"subtitle"`
	Body      string   `json:"body"`
	Tags      *Tags    `json:"tags"`
	Product   *Product `json:"product"`
	Published *Publish `json:"published"`
}

type Publish struct {
	Date   time.Time `json:"date"`
	Author string    `json:"author"` //uid
}

type Product struct {
	MainProduct *MainProductDetails `json:"main_product"`
	SubProducts []string            `json:"sub_products"`
	Details     string              `json:"details"`
}

type MainProductDetails struct {
	Name    string `json:"name"`
	Year    int    `json:"year"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type Tags struct {
	Tags []string `json:"tags"`
}

// API response
type ArticleResponse struct {
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Body        string    `json:"body"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Date        time.Time `json:"date"`
	AuthorName  string    `json:"author_name"`
	AuthorEmail string    `json:"author_email"`
}
