package models

import "time"

type User struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	UserName  string    `json:"user_name" valid:"required" validate:"required"`
	Email     string    `json:"email" valid:"required,email" gorm:"unique" validate:"required"`
	Password  string    `json:"password,omitempty" valid:"required" validate:"required"`
	Age       int       `json:"age" valid:"required" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLogin struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" valid:"required,email" gorm:"unique" validate:"required"`
	Password string `json:"password,omitempty" valid:"required" validate:"required"`
}

type UserUpdate struct {
	Email    string `json:"email" gorm:"unique" validate:"required"`
	UserName string `json:"user_name" validate:"required"`
}
