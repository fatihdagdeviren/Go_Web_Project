package Entities

import "github.com/jinzhu/gorm"

type UserProfile struct{
	gorm.Model
	UserID int
	User User
	Email string
	PhoneNumber string
	Description string
}
