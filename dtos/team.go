package dtos

type TeamInsert struct {
	Name      string `json:"name"`
	ClassID   string `json:"classId"`
	ProjectID string `json:"projectId"`
}

type Team struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
