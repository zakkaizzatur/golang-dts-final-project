package models

import "time"

// Comment represents the model for a comment
type Comment struct {
	Base
	UserID uint	`json:"user_id"`
	PhotoID uint `json:"photo_id"`
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~Comment message is required"`
	User User `json:"-"`
	Photo Photo `json:"-"`
}

type CommentGet struct {
	ID	uint `json:"id"`
	Message string `json:"message"`
	PhotoID uint `json:"photo_id"`
	UserID uint `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedAt *time.Time `json:"created_at"`
	User CommentUserGet `json:"user"`
	Photo CommentPhotoGet `json:"photo"`
}

type CommentPhotoGet struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Caption string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserID uint `json:"user_id"`
}

type CommentCreated struct {
	ID uint `json:"id"`
	Message string `json:"message"`
	PhotoID uint `json:"photo_id"`
	UserID uint `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
}

type CommentUserGet struct {
	ID uint `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type CommentUpdated struct {
	ID uint `json:"id"`
	Message string `json:"message"`
	UserID uint `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at"`
}
