package controllers

import (
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

// Register
func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindHeader(&User)
	} else {
		c.ShouldBind(&User)
	}

	passCheck := c.Request.FormValue("password")
	Check, _ := strconv.Atoi(passCheck)
	if Check < 6 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Password must be at least 6 characters long",
		})
	} else {
		err := db.Debug().Create(&User).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad request",
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"id":       User.ID,
			"email":    User.Email,
			"username": User.Username,
			"age":      User.Age,
		})
	}
}

// Login
func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindHeader(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Model(&User).Where("email = ?", User.Email).Take(&User).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))
	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}
	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
