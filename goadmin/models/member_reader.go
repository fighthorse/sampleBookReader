package models

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetMemberReaderTable(ctx *context.Context) table.Table {

	memberReader := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := memberReader.GetInfo()

	info.HideDetailButton()
	info.HideFilterButton()
	info.HideRowSelector()
	info.HideDeleteButton()
	info.SetFilterFormLayout(form.LayoutTwoCol)
	info.AddField("Id", "id", db.Int).
		FieldFilterable().
		FieldSortable()
	info.AddField("会员id", "member_id", db.Int).
		FieldFilterable().
		FieldSortable()
	info.AddField("书籍id", "book_id", db.Int).
		FieldFilterable().
		FieldSortable().
		FieldEditAble()
	info.AddField("章节id", "chapter_id", db.Int).
		FieldFilterable().
		FieldSortable().
		FieldEditAble()
	info.AddField("更新时间", "last_update", db.Varchar).
		FieldFilterable().
		FieldEditAble()

	info.SetTable("member_reader").SetTitle("会员书架").SetDescription("会员书架")

	formList := memberReader.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldHide()
	formList.AddField("会员id", "member_id", db.Int, form.Number)
	formList.AddField("书籍id", "book_id", db.Int, form.Number)
	formList.AddField("章节id", "chapter_id", db.Int, form.Number)
	formList.AddField("浏览时间", "last_update", db.Varchar, form.Text)

	formList.SetTable("member_reader").SetTitle("会员书架").SetDescription("会员书架")

	return memberReader
}
