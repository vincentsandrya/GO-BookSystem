package model

import "time"

type Buku struct {
	Id          int
	Title       string
	Description string
	ImageUrl    string `json:"image_url"`
	ReleaseYear int    `json:"release_year"`
	Price       int
	TotalPage   int `json:"total_page"`
	Thickness   string
	CategoryId  int `json:"category_id"`
	CreatedAt   time.Time
	CreatedBy   string `json:"created_by"`
	ModifiedAt  time.Time
	ModifiedBy  string `json:"modified_by"`
}

type Kategori struct {
	Id         int
	Name       string
	CreatedAt  time.Time
	CreatedBy  string `json:"created_by"`
	ModifiedAt time.Time
	ModifiedBy string `json:"modified_by"`
}
