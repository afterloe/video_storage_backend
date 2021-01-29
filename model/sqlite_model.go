package model

type Model struct {
	ModifyTime string `column:"modify_time" json:"modifyTime"` // 修改时间
	CreateTime string `column:"create_time" json:"createTime"` // 创建时间
	IsDel      bool   `column:"is_del" json:"isDel"`           // 是否删除
}

type User struct {
	Model
	ID       int64  `column:"id" json:"id"`             // 唯一标示
	Mail     string `column:"mail" json:"mail"`         // 邮箱
	Passwd   string `column:"pwd" json:"pwd"`           // 密码
	Nickname string `column:"nickname" json:"nickname"` // 昵称
	Avatar   string `column:"avatar" json:"avatar"`     // 头像
}

type Dictionary struct {
	Model
	ID   int64  `column:"id" json:"id"`     // 唯一标识
	Name string `column:"name" json:"name"` // 标签显示名称
	Data string `column:"data" json:"data"` // 标签值
}

type DictionaryGroup struct {
	Model
	ID     int64        `column:"id" json:"id"`     // 唯一标识
	Name   string       `column:"name" json:"name"` // 标签组显示名称
	Values []Dictionary `json:"values"`             // 标签组下所有的标签
}
