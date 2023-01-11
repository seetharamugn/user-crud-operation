package controllers

import (
	"go-crud/initializers"
	"go-crud/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// create a user
func UsersCreate(c *gin.Context) {

	var user models.User

	if temp, exists := c.Get("Fullname"); exists {
		user.Fullname = temp.(string)
	}
	if temp, exists := c.Get("Email"); exists {
		user.Email = temp.(string)
	}
	if temp, exists := c.Get("Mobile"); exists {
		user.Mobile = temp.(string)
	}
	if temp, exists := c.Get("Picture"); exists {
		user.Picture = temp.(string)
	}

	if result := initializers.DB.Create(&user); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":  "creaing user as problem ",
			"status": 500,
		})
		return
	}

	//Response with them
	c.JSON(http.StatusCreated, gin.H{
		"User":   user,
		"ststus": 200,
	})
}

// Fetch all the user with pagination
func FetchAll(c *gin.Context) {
	//Get the Users
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	var Users []models.User
	if result := initializers.DB.Limit(perPage).Offset((page - 1) * perPage).Find(&Users); result.Error != nil {
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
	var id int64 = 0
	if temp, exists := c.Get("userID"); exists {
		id = temp.(int64)
	}
	//Get the Users
	var User models.User
	if result := initializers.DB.First(&User, id); result.Error != nil {
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
	var user models.User
	var updateData models.User
	// Get the user ID and updateData from the context
	var id int64
	if temp, exists := c.Get("userID"); exists {
		id = temp.(int64)
	}
	if temp, exists := c.Get("updateUser"); exists {
		updateData = temp.(models.User)
	}

	if result := initializers.DB.First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found", "status": 404})
		return
	}
	//update it
	if result := initializers.DB.Model(&user).Updates(updateData); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":  "User not updated",
			"status": 500,
		})
		return
	}
	//Respond with it
	c.JSON(http.StatusOK, gin.H{
		"Users":  user,
		"ststus": 200,
	})
}

// delete the perticular user
func UsersDelete(c *gin.Context) {
	//Get the id off the url
	var id int64
	if temp, exists := c.Get("userID"); exists {
		id = temp.(int64)
	}
	//delete the User
	if result := initializers.DB.Delete(&models.User{}, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":  "User Not present in the DB",
			"status": 404,
		})
		return
	}

	//Respond with it
	c.Status(http.StatusOK)

}
