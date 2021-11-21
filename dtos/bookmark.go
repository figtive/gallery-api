package dtos

import "time"

type Bookmark struct {
	ID           string    `json:"id"`
	UserID       string    `json:"userId"`
	CourseworkID string    `json:"courseworkId"`
	CreatedAt    time.Time `json:"createdAt"`
}

type BookmarkAction struct {
	Mark bool `json:"mark"`
}
