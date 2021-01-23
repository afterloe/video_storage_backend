package model

// leechbox 安装应用信息
type AppTotal struct {
	Num        int   `json:"num"`        // 已安装的app数量
	RunTimeNum int   `json:"runTimeNum"` // 正在运行的app数量
	Apps       []App `json:"apps"`       // app信息
}

type App struct {
	Name        string `json:"name"`        // app 名称
	IsRunning   bool   `json:"isRunning"`   // 是否运行
	InstallPath string `json:"installPath"` // 安装位置
	Feature     string `json:"feature"`     // 功能描述
	Version     string `json:"version"`     // 版本
}

type Feature struct {
	Func []struct {
		Name     string `json:"name"`
		Describe string `json:"describe"`
		Command  string `json:"command"`
	} `json:"func"`
	AIP struct {
		Online []struct {
			Name       string      `json:"name"`
			Method     string      `json:"method"`
			URL        string      `json:"url"`
			Parameters []Parameter `json:"parameters"`
		} `json:"online"`

		Offline []struct {
			Parameter []Parameter `json:"parameter"`
		} `json:"offline"`
	} `json:"aip"`
	Demo []Demo `json:"demo"`
}

type Demo struct {
	Name       string      `json:"name"`
	Command    string      `json:"command"`
	Parameters []Parameter `json:"parameters"`
}

type Parameter struct {
	Name     string `json:"name"`
	Describe string `json:"describe"`
	Type     string `json:"type"`
}
