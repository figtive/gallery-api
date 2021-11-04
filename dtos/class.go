package dtos

type Class struct {
	ID          string `binding:"required" json:"id"`
	Name        string `binding:"required" json:"name"`
	Description string `binding:"required" json:"description"`
}
