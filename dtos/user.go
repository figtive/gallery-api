package dtos

type UserLogin struct {
	Token string `json:"token" binding:"required"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
