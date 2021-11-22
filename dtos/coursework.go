package dtos

import "time"

type Coursework struct {
	ID             string    `json:"id"`
	CourseID       string    `json:"courseId"`
	CreatedAt      time.Time `json:"createdAt"`
	CourseworkType string    `json:"courseworkType"`
}
