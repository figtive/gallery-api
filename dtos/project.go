package dtos

import "time"

type Project struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Active      bool      `json:"active"`
	Description string    `json:"description"`
	Field       string    `json:"field"`
	Thumbnail   string    `json:"thumbnail"`
	CreatedAt   time.Time `json:"createdAt"`
}

type ProjectInsert struct {
	Name        string `json:"name"`
	Active      bool   `json:"active"`
	Description string `json:"description"`
	Field       string `json:"field"`
	ClassID     string `json:"classId"`
}
