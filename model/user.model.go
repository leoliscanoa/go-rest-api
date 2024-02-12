package model

type User struct {
	ID        uint64 `json:"id" gorm:"primary_key:auto_increment"`
	Username  string `json:"username" gorm:"type:varchar(255) not null unique"`
	Firstname string `json:"firstname" gorm:"type:varchar(255) not null"`
	Lastname  string `json:"lastname" gorm:"type:varchar(255) not null"`
	Age       int    `json:"age" gorm:"type:int not null"`
}
