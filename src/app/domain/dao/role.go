package dao

type Role struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
	BaseModel
}
