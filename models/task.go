package models

type Task struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	AssignedTo string `json:"assignedTo"`
	Task       string `json:"task"`
}
