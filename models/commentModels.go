package models

import "time"

type Comment struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Message   string    `json:"message" validate:"required"`
	UserId    int       `json:"user_id"`
	PhotoId   int       `json:"photo_id"`
	User      *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	Photo     *Photo    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photo,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommentResponse struct {
	Id        int                     `json:"id" gorm:"primaryKey"`
	Message   string                  `json:"message" validate:"required"`
	UserId    int                     `json:"user_id"`
	PhotoId   int                     `json:"photo_id"`
	User      *UserCommentRelational  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	Photo     *PhotoCommentRelational `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photo,omitempty"`
	CreatedAt time.Time               `json:"created_at"`
	UpdatedAt time.Time               `json:"updated_at"`
}

type UserCommentRelational struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" validate:"required"`
	UserName string `json:"user_name" validate:"required"`
}

type PhotoCommentRelational struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" validate:"required"`
	UserId   int    `json:"user_id"`
}
