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

type Object struct {
	Model
	ID          int64  `column:"id" json:"id"`
	VirtualPath string `column:"virtual_path" json:"virtual_path"`
	MetadataID  int64  `column:"metadata_id" json:"metadata_id"` // 元数据id
}

type VideoDescribe struct {
	Model
	ID                 int64  `column:"id" json:"id"`                                     // 唯一标识
	MetadataID         int64  `column:"metadata_id" json:"metadata_id"`                   // 元数据id
	Width              int64  `column:"width" json:"width"`                               // 视频宽度
	Height             int64  `column:"height" json:"height"`                             // 视频高度
	Duration           string `column:"duration" json:"duration"`                         // 视频时间
	CodecName          string `column:"codec_name" json:"codec_name"`                     // 编码名称
	DisplayAspectRatio string `column:"display_aspect_ratio" json:"display_aspect_ratio"` // 播放比例
	CodecLongName      string `column:"codec_long_name" json:"codec_long_name"`           // 编码信息
}

type FileMetadata struct {
	ID         int64  `column:"id" json:"id"`                  // 唯一标识
	HexCode    string `column:"hex_code" json:"hexCode"`       // 唯一编码
	FileName   string `column:"filename" json:"fileName"`      // 源文件名称
	FileSize   int64  `column:"file_size" json:"fileSize"`     // 源文件大小
	FileType   string `column:"file_type" json:"file_type"`    // 源文件类型
	FullPath   string `column:"fullpath" json:"fullpath"`      // 源文件路径
	CreateTime string `column:"create_time" json:"createTime"` // 创建时间
	IsDel      bool   `column:"is_del" json:"isDel"`           // 是否删除
	IsLink     bool   `column:"is_link" json:"isLink"`         // 是否入库关联
}

type VideoDescribePackage struct {
	ID                 int64  `column:"id" json:"id"`               // 唯一标识
	FileName           string `column:"filename" json:"fileName"`   // 源文件名称
	FileType           string `column:"file_type" json:"file_type"` // 源文件类型
	FileSize           int64  `column:"file_size" json:"fileSize"`  // 源文件大小
	VirtualPath        string `column:"virtual_path" json:"virtual_path"`
	Duration           string `column:"duration" json:"duration"`                         // 视频时间
	CodecName          string `column:"codec_name" json:"codec_name"`                     // 编码名称
	Width              int64  `column:"width" json:"width"`                               // 视频宽度
	Height             int64  `column:"height" json:"height"`                             // 视频高度
	DisplayAspectRatio string `column:"display_aspect_ratio" json:"display_aspect_ratio"` // 播放比例
}
