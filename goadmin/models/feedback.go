package models

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetFeedbackTable(ctx *context.Context) table.Table {

	feedback := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := feedback.GetInfo()

	info.HideExportButton()

	info.HideDeleteButton()

	info.HideRowSelector()

	info.SetFilterFormLayout(form.LayoutTwoCol)
	info.AddField("Id", "id", db.Int)
	info.AddField("用户id", "member_id", db.Int).
		FieldFilterable().
		FieldSortable()
	info.AddField("反馈内容", "feed", db.Varchar)
	info.AddField("反馈时间", "feed_day", db.Varchar)
	info.AddField("回复内容", "callback", db.Varchar)
	info.AddField("回复人", "user_id", db.Int).
		FieldFilterable()

	info.SetTable("feedback").SetTitle("用户反馈记录").SetDescription("用户反馈记录")

	formList := feedback.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Member_id", "member_id", db.Int, form.Number).
		FieldDisableWhenUpdate()
	formList.AddField("Feed", "feed", db.Varchar, form.Text).
		FieldDisableWhenUpdate()
	formList.AddField("Feed_day", "feed_day", db.Varchar, form.Text).
		FieldDisableWhenUpdate()
	formList.AddField("Callback", "callback", db.Varchar, form.Text)
	formList.AddField("User_id", "user_id", db.Int, form.Number)

	formList.SetTable("feedback").SetTitle("用户反馈表").SetDescription("用户反馈表")

	return feedback
}
