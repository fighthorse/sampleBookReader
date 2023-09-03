package models

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetCommentTable(ctx *context.Context) table.Table {

	comment := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := comment.GetInfo()

	info.HideEditButton()

	info.SetFilterFormLayout(form.LayoutTwoCol)
	info.AddField("Id", "id", db.Int)
	info.AddField("书籍id", "book_id", db.Int).
		FieldFilterable().
		FieldSortable()
	info.AddField("章节id", "chapter_id", db.Int).
		FieldFilterable().
		FieldSortable()
	info.AddField("用户id", "member_id", db.Int).
		FieldFilterable().
		FieldSortable()
	info.AddField("评论", "comment_desc", db.Varchar)
	info.AddField("点赞数量", "comment_like", db.Int).
		FieldHide()

	info.SetTable("comment").SetTitle("书籍评论").SetDescription("书籍评论")

	formList := comment.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldHide()
	formList.AddField("书籍id", "book_id", db.Int, form.Number)
	formList.AddField("章节id", "chapter_id", db.Int, form.Number)
	formList.AddField("会员id", "member_id", db.Int, form.Number)
	formList.AddField("评论内容", "comment_desc", db.Varchar, form.Text)
	formList.AddField("点赞数量", "comment_like", db.Int, form.Number).
		FieldDefault("0").
		FieldHide()

	formList.SetTable("comment").SetTitle("书籍评论").SetDescription("Comment")

	return comment
}
