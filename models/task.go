package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	TaskName        string `json:"taskName"`
	TaskDescription string `json:"taskDescription"`
	TaskStatus      bool   `json:"taskStatus"`
	UserID          *uint
	User            *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
