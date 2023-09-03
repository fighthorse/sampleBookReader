package models

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetChapterTable(ctx *context.Context) table.Table {

	chapter := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := chapter.GetInfo()
	info.SetFilterFormLayout(form.LayoutTwoCol)
	info.AddField("章节id", "id", db.Int).
		FieldFilterable().
		FieldSortable()
	info.AddField("书籍id", "book_id", db.Int).
		FieldFilterable().
		FieldSortable().
		FieldEditAble()
	info.AddField("章节标题", "chapter_name", db.Varchar).
		FieldFilterable().
		FieldEditAble()
	info.AddField("章节内容", "chapter_content", db.Longtext).
		FieldHide()
	info.AddField("章节排序", "chapter_rank", db.Int).
		FieldFilterable().
		FieldSortable().
		FieldEditAble()

	info.SetTable("chapter").SetTitle("小说章节").SetDescription("小说章节列表")

	formList := chapter.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldHide()
	formList.AddField("书籍id", "book_id", db.Int, form.Number)
	formList.AddField("章节排序", "chapter_rank", db.Int, form.Number)
	formList.AddField("章节名称", "chapter_name", db.Varchar, form.Text)
	formList.AddField("章节内容", "chapter_content", db.Longtext, form.RichText)

	formList.SetTable("chapter").SetTitle("小说章节管理").SetDescription("小说章节")

	return chapter
}
