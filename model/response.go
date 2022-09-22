package model

import "os"

type ResponseBody struct {
	Code    int         `json:"code"`    // 返回码
	Message string      `json:"message"` // 错误信息
	Context interface{} `json:"context"` // 返回对象
}

type ContextPkg struct {
	Command    string      `json:"command"`    // 收到的命令
	Parameters interface{} `json:"parameters"` // 命令对应的参数
	Time       string      `json:"time"`       // 收到的信息的时间
}

type ScanFile struct {
	Name       string      `json:"name"`       // 文件名
	Path       string      `json:"path"`       // 文件全路径
	Size       int64       `json:"size"`       // 文件大小
	Mode       os.FileMode `json:"mode"`       // 文件类型
	ModifyTime string      `json:"modifyTime"` // 文件创建时间
}

type LoginBody struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
}

type ListBody struct {
	Data  interface{} `json:"data"`
	Total int         `json:"total"`
}
