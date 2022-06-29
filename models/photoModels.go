package models

import "time"

type Photo struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" validate:"required"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url" validate:"required"`
	UserId    int       `json:"user_id"`
	User      *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PhotoResponse struct {
	Id        int             `json:"id" gorm:"primaryKey"`
	Title     string          `json:"title" validate:"required"`
	Caption   string          `json:"caption"`
	PhotoURL  string          `json:"photo_url" validate:"required"`
	UserId    int             `json:"user_id"`
	User      *UserRelational `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type UserRelational struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" validate:"required"`
	UserName string `json:"user_name" validate:"required"`
}
