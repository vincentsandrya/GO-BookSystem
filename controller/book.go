package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vincentsandrya/GO-BookSystem/database"
	"github.com/vincentsandrya/GO-BookSystem/model"
	"github.com/vincentsandrya/GO-BookSystem/repository"
)

func GetCategories(c *gin.Context) {
	var (
		result gin.H
	)

	cat, err := repository.GetCategories(database.DbConnection)

	if err != nil {
		result = gin.H{
			"status":  "error",
			"message": err.Error(),
			"result":  nil,
		}

		c.JSON(http.StatusInternalServerError, result)
	} else {
		result = gin.H{
			"status":  "success",
			"message": "Success Getting Data!",
			"result":  cat,
		}

		c.JSON(http.StatusOK, result)
	}
}

func GetCategoryById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var (
		result gin.H
	)

	cat, err := repository.GetCategoryById(database.DbConnection, id)

	if cat == (model.Kategori{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Kategori tidak ditemukan!",
			"result":  nil,
		})
		panic("Kategori tidak ditemukan")
	}

	if err != nil {
		result = gin.H{
			"status":  "error",
			"message": err.Error(),
			"result":  nil,
		}

		c.JSON(http.StatusInternalServerError, result)
	} else {
		result = gin.H{
			"status":  "success",
			"message": "Success Getting Data!",
			"result":  cat,
		}

		c.JSON(http.StatusOK, result)
	}
}

func InsertCategories(c *gin.Context) {

	var cat model.Kategori

	err := c.BindJSON(&cat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
			"result":  nil,
		})
		panic(err)
	}

	checkExistsCat, _ := repository.GetCategoryByName(database.DbConnection, cat.Name)

	if checkExistsCat != (model.Kategori{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Nama kategori sudah tersedia!",
			"result":  nil,
		})
		panic("Nama kategori sudah tersedia!")
	}

	err = repository.InsertCategories(database.DbConnection, &cat)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
			"result":  nil,
		})
	} else {
		res, _ := repository.GetCategoryById(database.DbConnection, cat.Id)

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Success Insert Data!",
			"result":  res,
		})
	}
}

func DeleteCategories(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	checkExistsCat, _ := repository.GetCategoryById(database.DbConnection, id)

	if checkExistsCat == (model.Kategori{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Kategori tidak tersedia!",
			"result":  nil,
		})
		panic("Kategori tidak tersedia!")
	}

	err := repository.DeleteCategory(database.DbConnection, id)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
			"result":  nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Success Delete Data!",
			"result":  nil,
		})
	}
}

// BOOKS

func GetBooksByCategoryId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var (
		result gin.H
	)

	books, err := repository.GetBooksByCategoryId(database.DbConnection, id)

	if err != nil {
		result = gin.H{
			"status":  "error",
			"message": err.Error(),
			"result":  nil,
		}

		c.JSON(http.StatusInternalServerError, result)
	} else {
		result = gin.H{
			"status":  "success",
			"message": "Success Getting Data!",
			"result":  books,
		}

		c.JSON(http.StatusOK, result)
	}
}

func GetBooks(c *gin.Context) {
	var (
		result gin.H
	)

	books, err := repository.GetBooks(database.DbConnection)

	if err != nil {
		result = gin.H{
			"status":  "error",
			"message": err.Error(),
			"result":  nil,
		}

		c.JSON(http.StatusInternalServerError, result)
	} else {
		result = gin.H{
			"status":  "success",
			"message": "Success Getting Data!",
			"result":  books,
		}

		c.JSON(http.StatusOK, result)
	}
}

func GetBookById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var (
		result gin.H
	)

	buku, err := repository.GetBookById(database.DbConnection, id)

	if buku == (model.Buku{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Buku tidak ditemukan!",
			"result":  nil,
		})
		panic("Buku tidak ditemukan")
	}

	if err != nil {
		result = gin.H{
			"status":  "error",
			"message": err.Error(),
			"result":  nil,
		}

		c.JSON(http.StatusInternalServerError, result)
	} else {
		result = gin.H{
			"status":  "success",
			"message": "Success Getting Data!",
			"result":  buku,
		}

		c.JSON(http.StatusOK, result)
	}
}

func InsertBook(c *gin.Context) {

	var buku model.Buku

	err := c.BindJSON(&buku)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
			"result":  nil,
		})
		panic(err)
	}

	if buku.TotalPage > 100 {
		buku.Thickness = "tebal"
	} else {
		buku.Thickness = "tipis"
	}

	if buku.ReleaseYear < 1980 || buku.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Tahun rilis harus diantara 1980 sampai 2024!",
			"result":  nil,
		})
		panic("Tahun rilis harus diantara 1980 sampai 2024!")
	}

	err = repository.InsertBook(database.DbConnection, &buku)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
			"result":  nil,
		})
	} else {
		res, _ := repository.GetBookById(database.DbConnection, buku.Id)

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Success Insert Data!",
			"result":  res,
		})
	}
}

func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	checkExistsBook, _ := repository.GetBookById(database.DbConnection, id)

	if checkExistsBook == (model.Buku{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Buku tidak tersedia!",
			"result":  nil,
		})
		panic("Buku tidak tersedia!")
	}

	err := repository.DeleteBook(database.DbConnection, id)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
			"result":  nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Success Delete Data!",
			"result":  nil,
		})
	}
}
