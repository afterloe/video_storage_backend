package model

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
