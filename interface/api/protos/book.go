package protos

type PageSep struct {
	Ps int `json:"ps,omitempty" form:"ps"` // 单页数量
	Pn int `json:"pn,omitempty" form:"pn"` //页数
}

type MemberInfo struct {
	MemberName string `json:"member_name,omitempty" form:"member_name"`
	MemberId   string `json:"member_id,omitempty" form:"member_id"`
}
type CategoryReq struct {
	CategoryId int32 `json:"category_id,omitempty" form:"category_id"` //分类id
}

type BookListReq struct {
	Category int32  `json:"category_id,omitempty" form:"category_id"` //分类id
	Name     string `json:"name,omitempty" form:"name"`               //搜索名称
	Token    string `json:"token,omitempty" form:"token"`             //分类id
	PageSep
}

type BookListResp struct {
	Total int64       `json:"total,omitempty" form:"total"`
	Pn    int         `json:"pn,omitempty" form:"pn"`
	List  interface{} `json:"list,omitempty" form:"list"`
}

type MemberReq struct {
	CategoryId int32 `json:"category_id,omitempty" form:"category_id"` //分类id
	PageSep
	MemberInfo
}

type MemberBookListResp struct {
	Total int64       `json:"total,omitempty" form:"total"`
	Pn    int         `json:"pn,omitempty" form:"pn"`
	List  interface{} `json:"list,omitempty" form:"list"`
}

type AddShelfReq struct {
	BookId    int32 `json:"book_id,omitempty" form:"book_id"`
	ChapterId int32 `json:"chapter_id,omitempty" form:"chapter_id"`
	MemberInfo
}

type FeedBackReq struct {
	Content string `json:"content,omitempty" form:"content"` //分类id
	MemberInfo
}

type GetFeedBackReq struct {
	PageSep
	MemberInfo
}

type GetFeedBackResp struct {
	Total int64       `json:"total,omitempty" form:"total"`
	Pn    int         `json:"pn,omitempty" form:"pn"`
	List  interface{} `json:"list,omitempty" form:"list"`
}
