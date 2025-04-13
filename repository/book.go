package repository

import (
	"database/sql"

	"github.com/vincentsandrya/GO-BookSystem/model"
)

func GetCategories(db *sql.DB) (result []model.Kategori, err error) {
	sql := "SELECT * FROM Kategori"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var cat model.Kategori

		err = rows.Scan(&cat.Id, &cat.Name, &cat.CreatedAt, &cat.CreatedBy, &cat.ModifiedAt, &cat.ModifiedBy)
		if err != nil {
			return
		}

		result = append(result, cat)
	}

	return
}

func GetCategoryById(db *sql.DB, id int) (result model.Kategori, err error) {
	sql := "SELECT * FROM Kategori where id = $1"

	err = db.QueryRow(sql, id).Scan(
		&result.Id,
		&result.Name,
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

func GetCategoryByName(db *sql.DB, name string) (result model.Kategori, err error) {
	sql := "SELECT * FROM Kategori where name = $1"

	err = db.QueryRow(sql, name).Scan(
		&result.Id,
		&result.Name,
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

func InsertCategories(db *sql.DB, cat *model.Kategori) (err error) {
	sql := "INSERT INTO Kategori(name, created_by, modified_by) VALUES ($1, $2, $3) RETURNING id"

	errs := db.QueryRow(sql, cat.Name, cat.CreatedBy, cat.ModifiedBy).Scan(&cat.Id)

	if errs != nil {
		return errs
	}

	return nil
}

func DeleteCategory(db *sql.DB, id int) (err error) {
	sql := "DELETE FROM Kategori WHERE id = $1"

	errs := db.QueryRow(sql, id)
	return errs.Err()
}

// BOOKS

func GetBooksByCategoryId(db *sql.DB, catId int) (result model.Buku, err error) {
	sql := `SELECT B.* 
		FROM Kategori a
		INNER JOIN Buku b ON a.id = b.category_id
		WHERE a.id = $1`

	err = db.QueryRow(sql, catId).Scan(
		&result.Id,
		&result.Title,
		&result.Description,
		&result.ImageUrl,
		&result.ReleaseYear,
		&result.Price,
		&result.TotalPage,
		&result.Thickness,
		&result.CategoryId,
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

func GetBooks(db *sql.DB) (result []model.Buku, err error) {
	sql := "SELECT * FROM Buku"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var buku model.Buku

		err = rows.Scan(&buku.Id,
			&buku.Title,
			&buku.Description,
			&buku.ImageUrl,
			&buku.ReleaseYear,
			&buku.Price,
			&buku.TotalPage,
			&buku.Thickness,
			&buku.CategoryId,
			&buku.CreatedAt,
			&buku.CreatedBy,
			&buku.ModifiedAt,
			&buku.ModifiedBy)
		if err != nil {
			return
		}

		result = append(result, buku)
	}

	return
}

func GetBookById(db *sql.DB, id int) (res model.Buku, err error) {
	sql := "SELECT * FROM Buku where id = $1"

	err = db.QueryRow(sql, id).Scan(
		&res.Id,
		&res.Title,
		&res.Description,
		&res.ImageUrl,
		&res.ReleaseYear,
		&res.Price,
		&res.TotalPage,
		&res.Thickness,
		&res.CategoryId,
		&res.CreatedAt,
		&res.CreatedBy,
		&res.ModifiedAt,
		&res.ModifiedBy)

	if err != nil {
		return res, err
	}

	return res, nil
}

func InsertBook(db *sql.DB, buku *model.Buku) (err error) {
	sql := `INSERT INTO buku(
	title, description, image_url, release_year, price, total_page, thickness, category_id, created_by, modified_by) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`

	errs := db.QueryRow(sql,
		buku.Title,
		buku.Description,
		buku.ImageUrl,
		buku.ReleaseYear,
		buku.Price,
		buku.TotalPage,
		buku.Thickness,
		buku.CategoryId,
		buku.CreatedBy,
		buku.ModifiedBy).Scan(&buku.Id)

	if errs != nil {
		return errs
	}

	return nil
}

func DeleteBook(db *sql.DB, id int) (err error) {
	sql := "DELETE FROM buku WHERE id = $1"

	errs := db.QueryRow(sql, id)
	return errs.Err()
}
