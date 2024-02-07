package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        uint64 `json:"id,omitempty" gorm:"primary_key:auto_increment"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}
