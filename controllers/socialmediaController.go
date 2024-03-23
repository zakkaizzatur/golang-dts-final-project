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

func CreateSocialMedia(c *gin.Context){
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	parsedSocialMedia := parseResCreateSocialMedia(SocialMedia)

	c.JSON(http.StatusCreated, parsedSocialMedia)
}

func GetSocialMedia(c *gin.Context){
	db := database.GetDB()

	socialMedia := []models.SocialMedia{}

	err :=
		db.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "Email", "UserName")
		}).Find(&socialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	parsedSocialMedia := parseGetAllSocialMedia(socialMedia)

	c.JSON(http.StatusCreated, parsedSocialMedia)
}

func UpdateSocialMedia(c *gin.Context){
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialMediaId)

	err := db.Debug().Model(&SocialMedia).Where("id = ?", socialMediaId).Updates(models.SocialMedia{Name: SocialMedia.Name, SocialMediaUrl: SocialMedia.SocialMediaUrl}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	parsedSocialMedia := parseResUpdateSocialMedia(SocialMedia)

	c.JSON(http.StatusOK, parsedSocialMedia)
}

func DeleteSocialMedia(c *gin.Context){
	db := database.GetDB()
	
	SocialMedia := models.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))

	err := db.Model(&SocialMedia).Where("id = ?", socialMediaId).Delete(models.SocialMedia{}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}

func parseResCreateSocialMedia(socialMedia models.SocialMedia) models.SocialMediaCreated {
	parsedSocialMedia := models.SocialMediaCreated{
		ID: socialMedia.ID,
		Name: socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserID: socialMedia.UserID,
		CreatedAt: socialMedia.CreatedAt,
	}
	return parsedSocialMedia
	}

func parseResUpdateSocialMedia(socialMedia models.SocialMedia) models.SocialMediaUpdated{
	parsedSocialMedia := models.SocialMediaUpdated{
		ID: socialMedia.ID,
		Name: socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserID: socialMedia.UserID,
		UpdatedAt: socialMedia.UpdatedAt,
	}
	return parsedSocialMedia
	}

func parseGetAllSocialMedia(socialMedias []models.SocialMedia) []models.SocialMediaGet {
	var parsedSocialMedias []models.SocialMediaGet
	for _, socialMedia := range socialMedias {
		newSocialMedia := models.SocialMediaGet{
			ID: socialMedia.ID,
			Name: socialMedia.Name,
			SocialMediaUrl: socialMedia.SocialMediaUrl,
			UserID: socialMedia.UserID,
			User: models.PhotoUserGet{
				Email:    socialMedia.User.Email,
				Username: socialMedia.User.UserName,
			},
			CreatedAt: socialMedia.CreatedAt,
			UpdatedAt: socialMedia.UpdatedAt,
		}
		parsedSocialMedias = append(parsedSocialMedias, newSocialMedia)
	}
	return parsedSocialMedias
}