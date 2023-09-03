package models

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetMemberTable(ctx *context.Context) table.Table {

	member := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := member.GetInfo()
	info.HideExportButton()
	info.HideDeleteButton()
	info.SetFilterFormLayout(form.LayoutTwoCol)
	info.AddField("Id", "id", db.Int).
		FieldFilterable().
		FieldSortable()
	info.AddField("会员名称", "member_name", db.Varchar).
		FieldFilterable()
	info.AddField("密码", "member_pwd", db.Varchar).
		FieldHide()
	info.AddField("简介", "member_desc", db.Varchar)
	info.AddField("阅读数量", "read_books", db.Int)
	info.AddField("注册时间", "register_day", db.Varchar).
		FieldFilterable()

	info.SetTable("member").SetTitle("会员列表").SetDescription("会员列表")

	formList := member.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("会员名称", "member_name", db.Varchar, form.Text)
	formList.AddField("会员密码", "member_pwd", db.Varchar, form.Text)
	formList.AddField("简介", "member_desc", db.Varchar, form.Text)
	formList.AddField("阅读数量", "read_books", db.Int, form.Number)
	formList.AddField("注册时间", "register_day", db.Varchar, form.Text)

	formList.SetTable("member").SetTitle("书籍评论").SetDescription("Member")

	return member
}
