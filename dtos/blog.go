package dtos

type BlogInsert struct {
	Author   string `binding:"required" json:"author"`
	Title    string `binding:"required" json:"title"`
	Category string `binding:"required" json:"category"`
	Link     string `binding:"required" json:"link"`
	ClassID  string `binding:"required" json:"classId"`
}

type Blog struct {
	ID       string `json:"id"`
	Author   string `json:"author"`
	Title    string `json:"title"`
	Link     string `json:"link"`
	Category string `json:"category"`
}
