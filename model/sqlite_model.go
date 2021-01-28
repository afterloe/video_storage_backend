package model

type Model struct {
	ModifyTime string `column:"modify_time" json:"modifyTime"` // 修改时间
	CreateTime string `column:"create_time" json:"createTime"` // 创建时间
	IsDel      bool   `column:"is_del" json:"isDel"`           // 是否删除
}

type User struct {
	Model
	ID       int64  `column:"id" json:"id"` // 唯一标示
	Mail     string `column:"mail" json:"mail"`
	Passwd   string `column:"pwd" json:"pwd"`
	Nickname string `column:"nickname" json:"nickname"`
	Avatar   string `column:"avatar" json:"avatar"`
}
