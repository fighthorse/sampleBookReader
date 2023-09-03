package models

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetBookTable(ctx *context.Context) table.Table {

	book := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := book.GetInfo()

	info.HideDeleteButton()

	info.SetFilterFormLayout(form.LayoutTwoCol)
	info.AddField("Id", "id", db.Int).
		FieldSortable()
	info.AddField("书籍标题", "book_title", db.Varchar).
		FieldFilterable()
	info.AddField("简介", "book_desc", db.Varchar).
		FieldFilterable().
		FieldHide()
	info.AddField("作者id", "author_id", db.Int).
		FieldFilterable().
		FieldSortable()
	info.AddField("章节总数量", "chapter_total", db.Int).
		FieldFilterable()
	info.AddField("版权", "copyright", db.Varchar)
	info.AddField("状态", "state", db.Tinyint).
		FieldFilterable().
		FieldHide()
	info.AddField("分类", "category_id", db.Int).
		FieldFilterable().
		FieldSortable()

	info.SetTable("book").SetTitle("书籍管理").SetDescription("书籍管理")

	formList := book.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldHide()
	formList.AddField("书籍标题", "book_title", db.Varchar, form.Text)
	formList.AddField("简介", "book_desc", db.Varchar, form.Text)
	formList.AddField("作者id", "author_id", db.Int, form.Number)
	formList.AddField("章节数量", "chapter_total", db.Int, form.Number)
	formList.AddField("版权方", "copyright", db.Varchar, form.Text)
	formList.AddField("状态", "state", db.Tinyint, form.Number)
	formList.AddField("分类id", "category_id", db.Int, form.Number)

	formList.SetTable("book").SetTitle("书籍管理").SetDescription("书籍管理")

	return book
}
