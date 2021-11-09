package dtos

import "time"

type Coursework struct {
	ID        string    `json:"id"`
	CourseID  string    `json:"course_id"`
	CreatedAt time.Time `json:"created_at"`
}
