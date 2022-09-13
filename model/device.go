package model

// 运行时设备信息
type DeviceInfo struct {
	CPUs           int    `json:"CPUs"`
	PlatformFamily string `json:"platformFamily"`
	KernelArch     string `json:"kernelArch"`
}

// 处理器信息
type Processor struct {
	Cores     int32   `json:"cores"`
	Name      string  `json:"name"`
	Mhz       float64 `json:"Mhz"`
	CacheSize int32   `json:"cacheSize"`
	Microcode string  `json:"microcode"`
}

// 运行时设备信息
type RunTimeDeviceInfo struct {
	Name                string `json:"name"`
	Num                 string `json:"num"`
	DevelopmentLocation string `json:"developmentLocation"`
	AssociationTime     string `json:"associationTime"`
	Usage               string `json:"usage"`
	CPUUsedPercent      string `json:"CPUUsedPercent"`
	AgentID             string `json:"agentID"`
	MacID               string `json:"macID"`
	UID                 string `json:"uid"`
	NetInterface        []struct {
		MTU       int      `json:"mtu"`
		Name      string   `json:"name"`
		Flags     []string `json:"flags"`
		AddrGroup []string `json:"addrGroup"`
	} `json:"interface"`
	CPUs   []Processor `json:"CPUs"`
	Memory struct {
		Total       uint64  `json:"total"`
		Available   uint64  `json:"available"`
		Used        uint64  `json:"used"`
		UsedPercent float64 `json:"usedPercent"`
		Free        uint64  `json:"free"`
	} `json:"memory"`
	FS struct {
		Path        string  `json:"path"`
		FSType      string  `json:"FSType"`
		Total       uint64  `json:"total"`
		UsedPercent float64 `json:"usedPercent"`
		Free        uint64  `json:"free"`
		Used        uint64  `json:"used"`
	} `json:"FS"`
	System struct {
		Hostname        string `json:"hostname"`
		OS              string `json:"os"`
		Platform        string `json:"platform"`
		PlatformFamily  string `json:"platformFamily"`
		PlatformVersion string `json:"platformVersion"`
		KernelVersion   string `json:"kernelVersion"`
		KernelArch      string `json:"kernelArch"`
		HostId          string `json:"hostId"`
	} `json:"system"`
}
