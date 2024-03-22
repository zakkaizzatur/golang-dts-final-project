package controllers

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateComment(c *gin.Context){
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	parsedComment := parseResCreateComment(Comment)

	c.JSON(http.StatusCreated, parsedComment)
}

func GetComments(c *gin.Context){
	db := database.GetDB()

	comments := []models.Comment{}

	err :=
		db.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "Email", "UserName")
		}).Preload("Photo",func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "Title", "Caption", "PhotoUrl", "UserID")
		}).Find(&comments).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	parsedComments := parseGetAllComments(comments)

	c.JSON(http.StatusCreated, parsedComments)
}

func UpdateComment(c *gin.Context){
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	err := db.Debug().Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{Message: Comment.Message}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	parsedComment := parseResUpdateComment(Comment)

	c.JSON(http.StatusOK, parsedComment)
}

func DeleteComment(c *gin.Context){
	db := database.GetDB()
	
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))

	err := db.Model(&Comment).Where("id = ?", commentId).Delete(models.Comment{}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}

func parseResCreateComment(comment models.Comment) models.CommentCreated {
	parsedComment := models.CommentCreated{
		ID: comment.ID,
		Message: comment.Message,
		PhotoID: comment.PhotoID,
		UserID: comment.UserID,
		CreatedAt: comment.CreatedAt,
	}
	return parsedComment
	}

func parseResUpdateComment(comment models.Comment) models.CommentUpdated {
	parsedComment := models.CommentUpdated{
		ID: comment.ID,
		Message: comment.Message,
		UserID: comment.UserID,
		UpdatedAt: comment.UpdatedAt,
	}
	return parsedComment
	}

func parseGetAllComments(comments []models.Comment) []models.CommentGet {
	var parsedComments []models.CommentGet
	for _, comment := range comments {
		newComment := models.CommentGet{
			ID: comment.ID,
			Message:    comment.Message,
			PhotoID:  comment.PhotoID,
			UserID: comment.UserID,
			User: models.PhotoUserGet{
				Email:    comment.User.Email,
				Username: comment.User.UserName,
			},
			Photo: models.CommentPhotoGet{
				ID: comment.Photo.ID,
				Title: comment.Photo.Title,
				Caption: comment.Photo.Caption,
				PhotoUrl: comment.Photo.PhotoUrl,
				UserID: comment.Photo.UserID,
			},
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		}
		parsedComments = append(parsedComments, newComment)
	}
	return parsedComments
}