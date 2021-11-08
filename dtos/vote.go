package dtos

type VoteInsert struct {
	UserID       string `json:"userId"`
	CourseworkID string `json:"courseworkId"`
}
