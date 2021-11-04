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
	Name        string `binding:"required" json:"name"`
	Active      bool   `binding:"required" json:"active"`
	Description string `binding:"required" json:"description"`
	Field       string `binding:"required" json:"field"`
	ClassID     string `binding:"required" json:"classId"`
	Team        string `binding:"required" json:"team"`
	Metadata    string `binding:"required" json:"metadata"`
}

type ProjectThumbnail struct {
	ID string `form:"id"`
}
