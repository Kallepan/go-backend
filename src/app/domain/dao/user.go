package dao

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Status   int    `json:"status"`
	RoleID   int    `json:"role_id"`
	Role     Role   `json:"role"`
	BaseModel
}
