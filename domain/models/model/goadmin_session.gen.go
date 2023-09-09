// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameGoadminSession = "goadmin_session"

// GoadminSession mapped from table <goadmin_session>
type GoadminSession struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Sid       string    `gorm:"column:sid;not null" json:"sid"`
	Values    string    `gorm:"column:values;not null" json:"values"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName GoadminSession's table name
func (*GoadminSession) TableName() string {
	return TableNameGoadminSession
}
