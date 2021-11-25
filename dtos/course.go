package dtos

type Course struct {
	ID          string `binding:"required" json:"id"`
	Name        string `binding:"required" json:"name"`
	Description string `binding:"required" json:"description"`
	VoteQuota   int    `binding:"required" json:"voteQuota"`
}

type CourseUpdate struct {
	ID          string `binding:"required" json:"id"`
	Name        string `binding:"required" json:"name"`
	Description string `binding:"required" json:"description"`
	VoteQuota   int    `binding:"required" json:"voteQuota"`
}
