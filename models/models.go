package models

import "time"

type Todo struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}