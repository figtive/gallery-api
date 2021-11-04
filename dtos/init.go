package dtos

type Query struct {
	ID    string `form:"id"`
	Skip  int    `form:"skip"`
	Limit int    `form:"limit"`
}
