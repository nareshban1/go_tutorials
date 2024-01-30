package models

type Task struct {
	ID              uint   `json:"id" gorm:"primary_key"`
	TaskName        string `json:"taskName"`
	TaskDescription string `json:"taskDescription"`
	TaskStatus      bool   `json:"taskStatus"`
}
