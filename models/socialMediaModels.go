package models

import "time"

type SocialMedia struct {
	Id             int       `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name" validate:"required"`
	SocialMediaURL string    `json:"social_media_url" validate:"required"`
	UserId         int       `json:"user_id"`
	User           *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type SocialMediaResponse struct {
	Id             int                        `json:"id" gorm:"primaryKey"`
	Name           string                     `json:"name" validate:"required"`
	SocialMediaURL string                     `json:"social_media_url" validate:"required"`
	UserId         int                        `json:"user_id"`
	User           *UserSocialMediaRelational `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	CreatedAt      time.Time                  `json:"created_at"`
	UpdatedAt      time.Time                  `json:"updated_at"`
}

type UserSocialMediaRelational struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	UserName string `json:"user_name" validate:"required"`
}
