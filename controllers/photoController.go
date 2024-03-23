package controllers

import (
	"net/http"
	"strconv"

	"github.com/zakkaizzatur/golang-dts-final-project/database"
	"github.com/zakkaizzatur/golang-dts-final-project/helpers"
	"github.com/zakkaizzatur/golang-dts-final-project/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePhoto(c *gin.Context){
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	parsedPhoto := parseResCreatedPhoto(Photo)

	c.JSON(http.StatusCreated, parsedPhoto)
}

func GetPhotos(c *gin.Context){
	db := database.GetDB()

	photos := []models.Photo{}

	err :=
		db.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "Email", "UserName")
		}).Find(&photos).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	parsedPhotos := parseGetAllPhotos(photos)

	c.JSON(http.StatusCreated, parsedPhotos)
}

func UpdatePhoto(c *gin.Context){
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoUrl: Photo.PhotoUrl}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	parsedPhoto := parseResUpdatePhoto(Photo)

	c.JSON(http.StatusOK, parsedPhoto)
}

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))

	err := db.Model(&Photo).Where("id = ?", photoId).Delete(models.Photo{}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}


func parseGetAllPhotos(photos []models.Photo) []models.PhotoGet {
	var parsedPhotos []models.PhotoGet
	for _, photo := range photos {
		newPhoto := models.PhotoGet{
			ID:       photo.ID,
			Title:    photo.Title,
			Caption:  photo.Caption,
			PhotoUrl: photo.PhotoUrl,
			UserID:   photo.UserID,
			User: models.PhotoUserGet{
				Email:    photo.User.Email,
				Username: photo.User.UserName,
			},
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
		}
		parsedPhotos = append(parsedPhotos, newPhoto)
	}
	return parsedPhotos 
}

func parseResUpdatePhoto(photo models.Photo) models.PhotoUpdated {
	parsedPhoto := models.PhotoUpdated{
		ID: photo.ID,
		Title: photo.Title,
		Caption: photo.Caption,
		PhotoUrl: photo.PhotoUrl,
		UserID: photo.UserID,
		UpdatedAt: photo.UpdatedAt,
	}
	return parsedPhoto
	}

func parseResCreatedPhoto(photo models.Photo) models.PhotoCreated {
	parsedPhoto := models.PhotoCreated{
		ID: photo.ID,
		Title: photo.Title,
		Caption: photo.Caption,
		PhotoUrl: photo.PhotoUrl,
		UserID: photo.UserID,
		CreatedAt: photo.CreatedAt,
	}
	return parsedPhoto
	}