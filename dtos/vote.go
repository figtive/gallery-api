package dtos

type Vote struct {
	ID           string `json:"id"`
	UserID       string `json:"userId"`
	CourseworkID string `json:"courseworkId"`
}

type VoteInsert struct {
	CourseworkID string `json:"courseworkId"`
}

type VoteQuery struct {
	CourseWorkID string `json:"courseworkId"`
}
