package model

type User struct {
	ID        uint64 `json:"id" gorm:"primary_key:auto_increment"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}
