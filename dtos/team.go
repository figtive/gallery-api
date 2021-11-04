package dtos

type TeamInsert struct {
	Name      string `binding:"required" json:"name"`
	ClassID   string `binding:"required" json:"classId"`
	ProjectID string `binding:"required" json:"projectId"`
}

type Team struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
