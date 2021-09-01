package models

import "time"

type Books struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	AuthorID  int  `json:"-"`
	PublisherID  int  `json:"-"`
	Author    Authors         `json:"author"`
	Publisher   Publishers         `json:"publisher"`
	BookGenre   BookGenre `json:"-"`
	CreatedAt   time.Time    `json:"-"`
	UpdatedAt   time.Time    `json:"-"`
}

type Genre struct {
	ID       int    `json:"id"`
	GenreName string `json:"genre_name"`
	CreatedAt   time.Time    `json:"-"`
	UpdatedAt   time.Time    `json:"-"`
}

type BookGenre struct{
	ID int `json:"id"`
	BookID int `json:"-"`
	GenreID int `json:"-"`
	Genre Genre  `json:"genre"`
	CreatedAt   time.Time    `json:"-"`
	UpdatedAt   time.Time    `json:"-"`
	
}

type Authors struct {
	ID          int    `json:"-"`
	Name        string `json:"name"`
	Nationality string `json:"nationality"`
	CreatedAt   time.Time    `json:"-"`
	UpdatedAt   time.Time    `json:"-"`
}

type Publishers struct {
	ID   int    `json:"-"`
	Name string `json:"name"`
	PublisherID  int  `json:"-"`
	CreatedAt   time.Time    `json:"-"`
	UpdatedAt   time.Time    `json:"-"`
}