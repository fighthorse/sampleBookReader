package models

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetAuthorTable(ctx *context.Context) table.Table {

	author := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := author.GetInfo()

	info.SetFilterFormLayout(form.LayoutTwoCol)
	info.AddField("作者id", "id", db.Int).
		FieldFilterable()
	info.AddField("作者名称", "author_name", db.Varchar).
		FieldFilterable().
		FieldSortable()
	info.AddField("简介", "author_desc", db.Text).
		FieldHide()
	info.AddField("入住日期", "author_day", db.Varchar)
	info.AddField("性别", "author_sex", db.Varchar).
		FieldFilterable()

	info.SetTable("author").SetTitle("作者信息").SetDescription("作者信息")

	formList := author.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("Author_name", "author_name", db.Varchar, form.Text)
	formList.AddField("Author_desc", "author_desc", db.Text, form.RichText)
	formList.AddField("Author_day", "author_day", db.Varchar, form.Text)
	formList.AddField("Author_sex", "author_sex", db.Varchar, form.Text)

	formList.SetTable("author").SetTitle("作者信息").SetDescription("作者信息")

	return author
}
