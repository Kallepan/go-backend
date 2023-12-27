package dao

import (
	"github.com/google/uuid"
)

type Permission struct {
	// This is a simple permissions model. It is used to define what a user can do
	BaseModel

	Name        string  `gorm:"type:varchar(255);column:name;unique;not null" json:"name" binding:"required"`
	Description *string `gorm:"type:varchar(255);column:description" json:"description" binding:"omitempty"`
}

type Department struct {
	// This is a simple department model
	BaseModel

	Name string `gorm:"type:varchar(255);column:name;unique;not null" json:"name" binding:"required"`
}

type User struct {
	// This is a simple user model
	BaseModel

	Username string `gorm:"type:varchar(255);column:username;not null;index:idx_username,unique" json:"username" binding:"required,alpha,len=4,excludesall=!@#$%^&*()_+-="`
	// ->:false read-only field
	Password string `gorm:"type:varchar(255);column:password;->:false;<-:create" json:"-" binding:"required"`
	Email    string `gorm:"type:varchar(255);column:email;not null" json:"email" binding:"required,email"`

	// Each User belongs to a department
	DepartmentID uuid.UUID  `gorm:"type:uuid;column:department_id;not null" json:"department_id" binding:"required"`
	Department   Department `gorm:"foreignKey:DepartmentID;references:ID" json:"department" binding:"-"`

	// Each User has multiple permissions
	Permissions []Permission `gorm:"many2many:user_permissions;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"permissions"`
}