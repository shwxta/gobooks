package models

type Author struct {
	Name      string `json:"name"`
	BirthYear int    `json:"birthYear"`
}

type Publisher struct {
	Name        string `json:"name"`
	YearFounded int    `json:"yearFounded"`
}

type Book struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Author        Author    `json:"author"`
	Publisher     Publisher `json:"publisher"`
	Genres        []string  `json:"genres"`
	PublishedDate string    `json:"publishedDate"`
	Pages         int       `json:"pages"`
}
