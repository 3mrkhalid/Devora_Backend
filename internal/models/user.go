package models

import "time"

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username" binding:"required,min=3,max=30,alphanum"`
	Fullname  string    `json:"fullname" binding:"required,min=3,max=60"`
	Email     string    `json:"email" binding:"required,email,max=254"`
	Password  string    `json:"password" binding:"required,min=8,max=72"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}