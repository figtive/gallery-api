package dtos

type BlogInsert struct {
	Author   string `json:"author"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Link     string `json:"link"`
	ClassID  string `json:"classId"`
}

type Blog struct {
	ID       string `json:"id"`
	Author   string `json:"author"`
	Title    string `json:"title"`
	Link     string `json:"link"`
	Category string `json:"category"`
}
