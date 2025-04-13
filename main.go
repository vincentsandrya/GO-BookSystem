package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vincentsandrya/GO-BookSystem/controller"
	"github.com/vincentsandrya/GO-BookSystem/database"
	"github.com/vincentsandrya/GO-BookSystem/middleware"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Connect to PostgreSQL database
	database.ConnectDB()

	// Initialize Gin router
	r := gin.Default()

	// login doesnt need token
	v1 := r.Group("/api")
	v1.GET("/", defaultHandler)

	users := "/users"
	categories := "/categories"
	books := "/books"

	v1.POST(users+"/login", controller.Login)

	v1.GET(categories, middleware.AuthorizeHandlerCookies(), controller.GetCategories)
	v1.GET(categories+"/:id", middleware.AuthorizeHandlerCookies(), controller.GetCategoryById)
	v1.POST(categories, middleware.AuthorizeHandlerCookies(), controller.InsertCategories)
	v1.DELETE(categories+"/:id", middleware.AuthorizeHandlerCookies(), controller.DeleteCategories)
	v1.GET(categories+"/:id"+books, middleware.AuthorizeHandlerCookies(), controller.GetBooksByCategoryId)

	v1.GET(books, middleware.AuthorizeHandlerCookies(), controller.GetBooks)
	v1.GET(books+"/:id", middleware.AuthorizeHandlerCookies(), controller.GetBookById)
	v1.POST(books, middleware.AuthorizeHandlerCookies(), controller.InsertBook)
	v1.DELETE(books+"/:id", middleware.AuthorizeHandlerCookies(), controller.DeleteBook)

	// Start the server
	r.Run(":8080")
}

func defaultHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "DefaultGolang"})
}
