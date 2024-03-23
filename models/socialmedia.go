package models

import "time"

// Social Media represents the model for a social media
type SocialMedia struct {
	Base
	Name string `gorm:"not null" json:"name" form:"full_name" valid:"required~Social media name is required"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~Social media URL is required"`
	UserID uint
	User User `json:"-"`
}

type SocialMediaCreated struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserID uint `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
}

type SocialMediaUpdated struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserID uint `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type SocialMediaGet struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserID uint `json:"UserId"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	User PhotoUserGet `json:"User"`
}