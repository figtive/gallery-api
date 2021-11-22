package dtos

import (
	"mime/multipart"
	"time"
)

type Project struct {
	ID          string    `json:"id"`
	CourseID    string    `json:"courseId"`
	Name        string    `json:"name"`
	Team        string    `json:"team"`
	Description string    `json:"description"`
	Thumbnail   []string  `json:"thumbnail"`
	Link        string    `json:"link"`
	Video       string    `json:"video"`
	Field       string    `json:"field"`
	Active      bool      `json:"active"`
	Metadata    string    `json:"metadata"`
	CreatedAt   time.Time `json:"createdAt"`
}

type ProjectInsert struct {
	CourseID    string  `binding:"required" json:"courseId"`
	Name        string  `binding:"required" json:"name"`
	Team        string  `binding:"required" json:"team"`
	Description string  `binding:"required" json:"description"`
	Link        string  `binding:"required" json:"link"`
	Video       string  `binding:"required" json:"video"`
	Field       string  `binding:"required" json:"field"`
	Active      bool    `binding:"required" json:"active"`
	Metadata    *string `binding:"required" json:"metadata"`
}

type ProjectThumbnail struct {
	ID   string                `binding:"required" form:"id"`
	File *multipart.FileHeader `binding:"required" form:"file"`
}
