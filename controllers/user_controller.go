package controllers

import (
	"net/http"
	"time"

	"github.com/amos-pbi/final-task-pbi-rakamin-fullstack-amosaugusto/database"
	"github.com/amos-pbi/final-task-pbi-rakamin-fullstack-amosaugusto/middleware"
	"github.com/amos-pbi/final-task-pbi-rakamin-fullstack-amosaugusto/models"
	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var request struct {
		Username string
		Email    string
		Password string
	}

	if c.Bind(&request) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed request!",
		})
		return
	}

	if request.Username == "" || request.Email == "" || request.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username, email, or password cannot be empty!",
		})
		return
	}

	if len(request.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password's length minimum 6 characters!",
		})
		return
	}

	encrypt, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to encrypt the password!",
		})
		return
	}

	newUser := models.User{Username: request.Username, Email: request.Email, Password: string(encrypt)}
	res := database.DB.Create(&newUser)

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Registration Failed!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Registration Successful!",
	})
}

func Login(c *gin.Context) {
	var request struct {
		Email    string
		Password string
	}

	if c.Bind(&request) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed request",
		})
		return
	}

	var user models.User
	database.DB.First(&user, "email = ?", request.Email)
	if user.ID == 0 || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect Email or Password",
		})
		return
	}

	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":  user.ID,
		"expired": time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := sign.SignedString([]byte(middleware.SecretKey))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful!",
		"token":   token,
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully!",
	})
}

func UpdateUser(c *gin.Context) {
	user, _ := c.Get("user")
	id := c.Param("userId")

	var request struct {
		Username string
		Email    string
		Password string
	}

	if c.Bind(&request) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed request!",
		})
		return
	}

	if request.Username == "" || request.Email == "" || request.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username, email, or password cannot be empty!",
		})
		return
	}

	if len(request.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password's length minimum 6 characters!",
		})
		return
	}

	encrypt, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to encrypt the password!",
		})
		return
	}

	var updateUser models.User
	retrieve := database.DB.First(&updateUser, id)
	if retrieve.Error != nil || updateUser.ID != user.(models.User).ID {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found!",
		})
		return
	}

	jakartaLocation, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error loading Jakarta timezone!",
		})
		return
	}

	res := database.DB.Model(&updateUser).Updates(models.User{
		Username:   request.Username,
		Email:      request.Email,
		Password:   string(encrypt),
		Updated_At: time.Now().In(jakartaLocation).Format("2006-01-02 15:04:05"),
	})

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Update user failed!",
		})
		return
	}

	var updatedUser models.User
	database.DB.First(&updatedUser, updateUser.ID)

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":         updatedUser.ID,
			"username":   updatedUser.Username,
			"email":      updatedUser.Email,
			"created_at": updatedUser.Created_At,
			"updated_at": updatedUser.Updated_At,
		},
	})
}

func DeleteUser(c *gin.Context) {
	user, _ := c.Get("user")
	id := c.Param("userId")

	var deleteUser models.User
	retrieve := database.DB.First(&deleteUser, id)
	if retrieve.Error != nil || deleteUser.ID != user.(models.User).ID {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found!",
		})
		return
	}

	var photo models.Photo
	resPhoto := database.DB.Where("user_id = ?", user.(models.User).ID).Delete(&photo)
	if resPhoto.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Delete user's photos failed!",
		})
		return
	}

	res := database.DB.Delete(&deleteUser)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Delete user failed!",
		})
		return
	}

	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted and logged out successfully!",
	})
}
