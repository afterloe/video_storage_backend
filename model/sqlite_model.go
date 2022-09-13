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
	ID      int64  `column:"id" json:"id"`            // 唯一标识
	Name    string `column:"name" json:"name"`        // 标签显示名称
	Data    string `column:"data" json:"data"`        // 标签值
	GroupID int64  `column:"group_id" json:"groupID"` // 所属标签组
}

type DictionaryGroup struct {
	Model
	ID        int64         `column:"id" json:"id"`                // 唯一标识
	Name      string        `column:"name" json:"name"`            // 标签组显示名称
	GroupType string        `column:"group_type" json:"groupType"` // 标签组类型
	Values    []*Dictionary `json:"values"`                        // 标签组下所有的标签
}

type DemandVideo struct {
	Model
	ID         int64  `column:"id" json:"id"`                  // 唯一标识
	Name       string `column:"name" json:"name"`              // 视频名称
	Path       string `column:"path" json:"path"`              // 文件全路径
	Size       int64  `column:"size" json:"size"`              // 文件大小
	Width      int    `column:"width" json:"width"`            // 视频宽度
	Height     int    `column:"height" json:"height"`          // 视频高度
	Duration   float64    `column:"duration" json:"duration"`      // 视频时间
	Describe   string `column:"describe" json:"describe"`      // 描述
	Title      string `column:"title" json:"title"`            // 视频标题
	FFmpegJSON string `column:"ffmpeg_json" json:"FFmpegJSON"` // 视频流信息
}
