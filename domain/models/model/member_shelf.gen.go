// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMemberShelf = "member_shelf"

// MemberShelf mapped from table <member_shelf>
type MemberShelf struct {
	ID        int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	BookID    int32  `gorm:"column:book_id;not null" json:"book_id"`
	ChapterID int32  `gorm:"column:chapter_id;not null" json:"chapter_id"`
	ReadDay   string `gorm:"column:read_day;comment:浏览时间" json:"read_day"` // 浏览时间
	MemberID  int32  `gorm:"column:member_id;not null" json:"member_id"`
}

// TableName MemberShelf's table name
func (*MemberShelf) TableName() string {
	return TableNameMemberShelf
}