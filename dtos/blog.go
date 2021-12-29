package dtos

import "time"

type Blog struct {
	ID        string    `json:"id"`
	CourseID  string    `json:"courseId"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Link      string    `json:"link"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"createdAt"`
}

type BlogInsert struct {
	CourseID string `binding:"required" json:"courseId"`
	Title    string `binding:"required" json:"title"`
	Author   string `binding:"required" json:"author"`
	Link     string `binding:"required" json:"link"`
	Category string `binding:"required" json:"category"`
}

type BlogUpdate struct {
	ID       string `json:"id"`
	CourseID string `json:"courseId"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Link     string `json:"link"`
	Category string `json:"category"`
}

type BlogQuery struct {
	CourseworkQuery
	Title    string `form:"title"`
	Category string `form:"category"`
}
