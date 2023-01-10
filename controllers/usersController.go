package controllers

import (
	"go-crud/initializers"
	"go-crud/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// create a user
func UsersCreate(c *gin.Context) {
	//Get the User body from the request

	var body struct {
		Fullname string `validate:"required"`
		Email    string `validate:"required,email"`
		Mobile   int64  `validate:"required"`
		Picture  string `validate:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	User := models.User{
		Fullname: body.Fullname,
		Email:    body.Email,
		Mobile:   body.Mobile,
		Picture:  body.Picture,
	}

	result := initializers.DB.Create(&User)

	//handling error
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid request body",
			"status": 400,
		})
		return
	}

	//Response with them
	c.JSON(http.StatusOK, gin.H{
		"User":   User,
		"ststus": 200,
	})
}

// Fetch all the user with pagination
func FetchAll(c *gin.Context) {
	//Get the Users
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	var Users []models.User
	result := initializers.DB.Limit(perPage).Offset((page - 1) * perPage).Find(&Users)

	//handling error
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":  "Fetching users from the DB has problem ",
			"status": 500,
		})
		return
	}
	//Response with them
	c.JSON(http.StatusOK, gin.H{
		"Users": Users,
		"page":  page,
	})

}

// Fetch the perticular user
func FindByID(c *gin.Context) {
	//Get if off url
	id := c.Param("id")
	//Get the Users
	var User models.User
	result := initializers.DB.First(&User, id)

	//handling error
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":  "User Not Presnt in the Db",
			"status": 404,
		})
		return
	}
	//Response with them
	c.JSON(http.StatusOK, gin.H{
		"Users":  User,
		"ststus": 200,
	})
}

// update the perticualr user
func UsersUpdate(c *gin.Context) {
	//Get the id off the url\

	id := c.Param("id")
	//Find the User were to update
	var body struct {
		Fullname string
		Email    string
		Mobile   int64
		Picture  string
	}
	c.Bind(&body)

	//FEtching the data from the db
	var User models.User
	result1 := initializers.DB.First(&User, id)
	//handling error
	if result1.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":  "User Not present in the DB",
			"status": 404,
		})
		return
	}

	//update it
	result := initializers.DB.Model(&User).Updates(models.User{
		Fullname: body.Fullname,
		Email:    body.Email,
		Mobile:   body.Mobile,
		Picture:  body.Picture,
	})
	//handling error
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":  "User not updated",
			"status": 500,
		})
		return
	}
	//Respond with it
	c.JSON(http.StatusOK, gin.H{
		"Users":  User,
		"ststus": 200,
	})
}

// delete the perticular user
func UsersDelete(c *gin.Context) {
	//Get the id off the url
	id := c.Param("id")

	//delete the User
	result := initializers.DB.Delete(&models.User{}, id)

	//handling error
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":  "User Not present in the DB",
			"status": 404,
		})
		return
	}

	//Respond with it
	c.Status(http.StatusOK)

}
