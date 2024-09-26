package models

import "time"

type User struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Role     string `json:"role,omitempty" db:"role"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type Note struct {
	ID        int        `json:"id" db:"id"`
	Text      string     `json:"text" db:"text"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UserId    int        `json:"user_id" db:"user_id"`
	Deadline  *time.Time `json:"deadline" db:"deadline"`
}

type DeleteNoteRequest struct {
	ID int `json:"id" db:"id"`
}
