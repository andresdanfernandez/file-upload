package models

import "time"

type File struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	Filename   string    `json:"filename"`
	URL        string    `json:"url"`
	Size       int64     `json:"size"`
	UploadedAt time.Time `json:"uploaded_at"`
}
