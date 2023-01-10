package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Fullname string
	Email    string
	Mobile   int64
	Picture  string
}
