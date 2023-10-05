package models

type Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type Register struct {
	Email    string `json:"email" binding:"required,min=6,max=30"`
	Name     string `json:"name" binding:"required,max=255"`
	Password string `json:"password" binding:"required,max=255"`
}