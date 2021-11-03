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
	Metadata    string    `json:"metadata"`
	Team        string    `json:"team"`
}

type ProjectInsert struct {
	Name        string `json:"name"`
	Active      bool   `json:"active"`
	Description string `json:"description"`
	Field       string `json:"field"`
	ClassID     string `json:"classId"`
	Team        string `json:"team"`
	Metadata    string `json:"metadata"`
}

type ProjectThumbnail struct {
	ID string `form:"id"`
}
