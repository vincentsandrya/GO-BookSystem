package model

import "time"

type User struct {
	Id         int
	Username   string
	Password   string
	CreatedAt  time.Time
	CreatedBy  string `json:"created_by"`
	ModifiedAt time.Time
	ModifiedBy string `json:"modified_by"`
}
