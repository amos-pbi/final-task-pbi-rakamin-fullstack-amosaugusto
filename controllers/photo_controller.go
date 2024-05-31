package controllers

import (
	"net/http"
	"time"

	"github.com/amos-pbi/final-task-pbi-rakamin-fullstack-amosaugusto/database"
	"github.com/amos-pbi/final-task-pbi-rakamin-fullstack-amosaugusto/models"

	"github.com/gin-gonic/gin"
)

func ShowPhoto(c *gin.Context) {
	user, _ := c.Get("user")

	var photos []models.Photo
	database.DB.Where("user_id = ?", user.(models.User).ID).Find(&photos)

	c.JSON(http.StatusOK, gin.H{
		"photos": photos,
	})
}

func CreatePhoto(c *gin.Context) {
	user, _ := c.Get("user")

	var request struct {
		Title    string
		Caption  string
		PhotoUrl string
	}

	if c.Bind(&request) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed request!",
		})
		return
	}

	photo := models.Photo{
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoUrl: request.PhotoUrl,
		UserID:   user.(models.User).ID,
	}
	res := database.DB.Create(&photo)

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Create photo failed!",
		})
		return
	}

	var createdPhoto models.Photo
	database.DB.First(&createdPhoto, photo.ID)

	c.JSON(http.StatusOK, gin.H{
		"photo": createdPhoto,
	})
}

func UpdatePhoto(c *gin.Context) {
	user, _ := c.Get("user")
	id := c.Param("photoId")

	var request struct {
		Title    string
		Caption  string
		PhotoUrl string
	}

	if c.Bind(&request) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed request!",
		})
		return
	}

	var photo models.Photo
	retrieve := database.DB.First(&photo, id)
	if retrieve.Error != nil || photo.UserID != user.(models.User).ID {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Photo not found!",
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

	res := database.DB.Model(&photo).Updates(models.Photo{
		Title:      request.Title,
		Caption:    request.Caption,
		PhotoUrl:   request.PhotoUrl,
		Updated_At: time.Now().In(jakartaLocation).Format("2006-01-02 15:04:05"),
	})

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Update photo failed!",
		})
		return
	}

	var updatedPhoto models.Photo
	database.DB.First(&updatedPhoto, photo.ID)

	c.JSON(http.StatusOK, gin.H{
		"photo": updatedPhoto,
	})
}

func DeletePhoto(c *gin.Context) {
	user, _ := c.Get("user")
	id := c.Param("photoId")

	var photo models.Photo
	retrieve := database.DB.First(&photo, id)
	if retrieve.Error != nil || photo.UserID != user.(models.User).ID {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Photo not found!",
		})
		return
	}

	res := database.DB.Delete(&photo)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Delete photo failed!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success! Photo has been deleted",
	})
}
