package middleware

import (
	"fmt"
	"go-crud/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewUserMiddleware(c *gin.Context) {
	fullname := c.PostForm("Fullname")
	email := c.PostForm("Email")
	mobile := c.PostForm("Mobile")
	file, err := c.FormFile("file")
	fmt.Println(file.Filename)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":  "image is required ",
			"status": 400,
		})
		return
	}
	// Verify the input parameters
	if fullname == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "fullname is required"})
		c.Abort()
		return
	}

	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
		c.Abort()
		return
	}

	if mobile == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "mobile number is required"})
		c.Abort()
		return
	}
	picture := file.Filename
	fmt.Println(picture)
	c.SaveUploadedFile(file, "assets/uploads/"+file.Filename)
	// Set the user struct to the context so it can be accessed in the createUser function
	c.Set("Fullname", fullname)
	c.Set("Email", email)
	c.Set("Mobile", mobile)
	c.Set("Picture", picture)

	c.Next()
}

func GetUserMiddleware(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user ID is required"})
		c.Abort()
		return
	}
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user ID is invalid"})
		c.Abort()
		return
	}
	c.Set("userID", id)
	c.Next()
}

func UpdateUserMiddleware(c *gin.Context) {
	userID := c.Param("id")
	var updateUser models.User

	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user ID is required"})
		c.Abort()
		return
	}
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user ID is invalid"})
		c.Abort()
		return
	}

	// Bind JSON input to the user struct

	c.Set("userID", id)
	c.Set("updateUser", updateUser)
	c.Next()
}
