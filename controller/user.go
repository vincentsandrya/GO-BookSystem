package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vincentsandrya/GO-BookSystem/database"
	"github.com/vincentsandrya/GO-BookSystem/dto"
	"github.com/vincentsandrya/GO-BookSystem/model"
	"github.com/vincentsandrya/GO-BookSystem/repository"
)

func Login(c *gin.Context) {

	var user model.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
			"result":  nil,
		})
		panic(err)
	}

	user, err = repository.Login(database.DbConnection, user.Username, user.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
			"result":  nil,
		})
		panic(err)
	}

	if user == (model.User{}) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Username atau password salah!",
			"result":  nil,
		})
		panic("username atau password salah!")
	}

	// token, err := dto.GenerateAccessToken(dto.JWTClaims{
	// 	ID:       int(user.Id),
	// 	Username: user.Username,
	// })

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"status":  "error",
	// 		"message": err.Error(),
	// 		"result":  nil,
	// 	})
	// 	panic(err)
	// }

	// Generate JWT token on successful login
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &dto.JWTClaims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	jwtKey := []byte(os.Getenv("API_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("Could not create token", http.StatusInternalServerError))
		return
	}

	// Set token as HttpOnly cookie
	c.SetCookie(
		"authToken",
		tokenString,
		int(expirationTime.Unix()-time.Now().Unix()),
		"/",
		"",   // Domain, can be set to your specific domain
		true, // Secure, set to true in production
		true, // HttpOnly
	)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Login berhasil!",
		"token":   tokenString,
		"result":  user,
	})

}
