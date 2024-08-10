package models

import "time"

type Task struct {
	ID          string    `json:"id" bson:"id" unique:"true"` 
	Title       string    `json:"title" bson:"title" required:"true"`
	Description string    `json:"description" bson:"description"`
	DueDate     time.Time `json:"due_date" bson:"due_date"`
	Status      string    `json:"status" bson:"status"`
}
