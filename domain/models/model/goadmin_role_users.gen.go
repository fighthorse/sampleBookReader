// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameGoadminRoleUser = "goadmin_role_users"

// GoadminRoleUser mapped from table <goadmin_role_users>
type GoadminRoleUser struct {
	RoleID    int32     `gorm:"column:role_id;primaryKey" json:"role_id"`
	UserID    int32     `gorm:"column:user_id;primaryKey" json:"user_id"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName GoadminRoleUser's table name
func (*GoadminRoleUser) TableName() string {
	return TableNameGoadminRoleUser
}