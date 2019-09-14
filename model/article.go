package model

import "time"

type Article struct {
	uid       string
	title     string
	subtitle  string
	body      string
	tags      *tags
	product   *product
	city      string
	country   string
	published publish
}

type publish struct {
	date   time.Time
	author string // uid
}

type product struct {
	mainProduct *productDetails
	subProducts []string
	details     string
}

type productDetails struct {
	name string
	year int
}

type tags struct {
	tag []string
}

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
