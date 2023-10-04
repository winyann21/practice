package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username	string	`gorm:"size:255;not null;unique" json:"username"`
	Password	string	`gorm:"size:255;not null;" json:"-"` //json "-" so it will not return on the response.
	Entries []Entry
}