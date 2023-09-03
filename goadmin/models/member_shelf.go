package models

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetMemberShelfTable(ctx *context.Context) table.Table {

	memberShelf := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := memberShelf.GetInfo()
	info.HideDeleteButton()
	info.SetFilterFormLayout(form.LayoutTwoCol)
	info.AddField("Id", "id", db.Int).
		FieldHide()
	info.AddField("会员id", "member_id", db.Int).
		FieldFilterable().
		FieldSortable()
	info.AddField("书籍id", "book_id", db.Int).
		FieldFilterable().
		FieldSortable()
	info.AddField("章节id", "chapter_id", db.Int).
		FieldFilterable().
		FieldSortable()
	info.AddField("浏览时间", "read_day", db.Varchar)

	info.SetTable("member_shelf").SetTitle("阅读记录").SetDescription("阅读记录")

	formList := memberShelf.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldHide()
	formList.AddField("会员id", "member_id", db.Int, form.Number)
	formList.AddField("书籍id", "book_id", db.Int, form.Number)
	formList.AddField("章节id", "chapter_id", db.Int, form.Number)
	formList.AddField("浏览时间", "read_day", db.Varchar, form.Text)

	formList.SetTable("member_shelf").SetTitle("会员书架").SetDescription("会员书架")

	return memberShelf
}
