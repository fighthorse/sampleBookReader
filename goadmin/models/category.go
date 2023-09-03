package models

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetCategoryTable(ctx *context.Context) table.Table {

	category := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := category.GetInfo()

	info.HideExportButton()

	info.HideDetailButton()
	info.SetFilterFormLayout(form.LayoutTwoCol)
	info.AddField("序号", "id", db.Int).
		FieldSortable()
	info.AddField("分类名称", "category_name", db.Varchar).
		FieldFilterable().
		FieldEditAble()
	info.AddField("父分类", "parent_id", db.Int).
		FieldFilterable().
		FieldSortable().
		FieldEditAble()
	info.AddField("是否删除", "is_del", db.Tinyint).
		FieldFilterable().
		FieldEditAble()
	info.SetTable("category").SetTitle("分类管理").SetDescription("分类管理")

	formList := category.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("分类名称", "category_name", db.Varchar, form.Text)
	formList.AddField("继承父分类id", "parent_id", db.Int, form.Number)
	formList.AddField("是否删除", "is_del", db.Tinyint, form.Number).
		FieldDefault("0")

	formList.SetTable("category").SetTitle("Category").SetDescription("Category")

	return category
}
