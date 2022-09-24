package config

// 应用程序配置
type ApplicationConfig struct {
	Common struct {
		ServerName string `ini:"server_name"` // 应用名
		Version    string `ini:"version"`     // 应用版本
		LogPath    string `ini:"log_path"`    // 日志文件输出目录
		SocketFile string `ini:"socket_file"` // socket监听文件全路径
		LogLevel   string `ini:"log_level"`   // 日志级别
	} `ini:"common"` // 通用系统配置

	Backend struct {
		ServiceHost string `ini:"server_host"` // 服务监听地址
		Port        string `ini:"port"`        // 监听端口
	} `ini:"backend"` // 守护进程服务配置

	Logic struct {
		SoftwareWarehouse string `ini:"software_warehouse"` // 软件仓库位置
		SQLiteFile        string `ini:"sqlite_file"`        // sqlite3 文件地址
		VideoStorage      string `ini:"video_storage"`      // video存放的位置
		Screenshot        string `ini:"screenshot"`         // 截图存放位置
		VideoPrefix       string `ini:"video_prefix"`       // 点播视频前缀
	} `ini:"logic"`
}
