package controllers

import "github.com/gin-gonic/gin"

func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func hashPassword(password string) string {
	return ""
}

func verifyPassword(userPassword string, providedPassword string) bool {
	return false
}
