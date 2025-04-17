package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	UUID       uuid.UUID `gorm:"column:id, type:int; not null; primaryKey; autoIncrement;"`
	RoleName string    `gorm:"column:role_name"`
	RoleNote string    `gorm:"column:role_note; type: text"`
}

func (r *Role) TableName() string {
	return "go_db_roles"
}
