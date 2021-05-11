package models

type Book struct {
	ID          string  `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	ReleaseDate string  `json:"releaseDate,omitempty"`
	Author      *Author `json:"author,omitempty"`
}

type Author struct {
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}