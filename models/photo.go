package models

import "time"

// Photo represents the model for a photo
type Photo struct {
	Base
	Title string `gorm:"not null" json:"title" form:"title" valid:"required~Title for the photo is required"`
	Caption string `gorm:"not null" json:"caption" form:"caption"`
	PhotoUrl string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~URL for the photo is required"`
	UserID uint `json:"user_id"`
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	User User `json:"-"`
}

type PhotoGet struct {
	ID        uint 			`json:"id"`
	Title     string       `json:"title"`
	Caption   string       `json:"caption"`
	PhotoUrl  string       `json:"photo_url"`
	UserID    uint          `json:"user_id"`
	User      PhotoUserGet `json:"user"`
	CreatedAt *time.Time    `json:"created_at"`
	UpdatedAt *time.Time    `json:"updated_at"`
}

type PhotoUserGet struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PhotoUpdated struct {
	ID        uint 			`json:"id"`
	Title     string       `json:"title"`
	Caption   string       `json:"caption"`
	PhotoUrl  string       `json:"photo_url"`
	UserID    uint          `json:"user_id"`
	UpdatedAt *time.Time    `json:"updated_at"`
}

type PhotoCreated struct {
	ID        uint 			`json:"id"`
	Title     string       `json:"title"`
	Caption   string       `json:"caption"`
	PhotoUrl  string       `json:"photo_url"`
	UserID    uint          `json:"user_id"`
	CreatedAt *time.Time    `json:"created_at"`
}