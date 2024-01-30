package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}
