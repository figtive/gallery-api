package dtos

type Query struct {
	ID      string `form:"id"`
	Skip    int    `form:"skip"`
	Limit   int    `form:"limit"`
	Current bool   `form:"current"`
}

type CourseworkQuery struct {
	Query
	CourseID string `uri:"course_id"`
}
