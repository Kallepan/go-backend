package dao

import (
	"time"
)

type BaseModel struct {
	// Add this model to all other structs to add common fields
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
