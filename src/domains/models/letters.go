package models

import "time"

// Letters struct
type Letters struct {
	ID        int       `json:"id"`
	To        string    `json:"to"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ReqCreateLetters struct
type ReqCreateLetters struct {
	To   string `json:"to" form:"to" binding:"required"`
	Text string `json:"text" form:"text" binding:"required"`
}

// ResCreateLetters struct
type ResCreateLetters struct {
	ID        int       `json:"id"`
	To        string    `json:"to"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
