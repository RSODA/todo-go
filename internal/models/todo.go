package models

import "time"

type TODO struct {
	ID        int64     `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Tasks     []Task    `json:"task" db:"task"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Task struct {
	ID          int64  `json:"id" db:"id"`
	TodoID      int64  `json:"todo_id" db:"todo_id"`
	Description string `json:"description" db:"description"`
	IsComplete  bool   `json:"is_complete" db:"is_complete"`
}
