package repository

import (
	"database/sql"

	"github.com/vincentsandrya/GO-BookSystem/model"
)

func Login(db *sql.DB, username string, password string) (result model.User, err error) {
	sql := "SELECT id, username, created_at, created_by, modified_at, modified_by FROM Users where username = $1 and password = $2"

	err = db.QueryRow(sql, username, password).Scan(
		&result.Id,
		&result.Username,
		&result.CreatedAt,
		&result.CreatedBy,
		&result.ModifiedAt,
		&result.ModifiedBy,
	)
	if err != nil {
		return result, err
	}

	return result, nil
}
