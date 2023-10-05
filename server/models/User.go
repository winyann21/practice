package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email		string	`gorm:"size:30;not null;unique" json:"email"`
	Name		string	`gorm:"size:255;not null;" json:"name"`
	Password	string	`gorm:"size:255;not null;" json:"-"` //json "-" so it will not return on the response.
	Entries []Entry
}