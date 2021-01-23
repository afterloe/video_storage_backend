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
		ServiceHost   string `ini:"server_host"`    // 服务监听地址
		Port          string `ini:"port"`           // 监听端口
		AMQPHost      string `ini:"amqp_host"`      // mq 地址
		AMQPPort      string `ini:"amqp_port"`      // mq 端口
		AMQPTimeout   string `ini:"amqp_timeout"`   // mq 超时
		AMQPUserName  string `ini:"amqp_user_name"` // mq 用户名
		AMQPPassword  string `ini:"amqp_password"`  // mq 密码
		AMQPRouting   string `ini:"routing_key"`    // mq routing key
		QueueVHost    string `ini:"queue_vhost"`    // mq vhost
		QueueExchange string `ini:"queue_exchange"` // mq exchange
	} `ini:"backend"` // 守护进程服务配置

	Ekho struct {
		VoiceTmpPath string `ini:"voice_tmp_path"` // 生成的语音存放位置
		VoiceSpeed   string `ini:"voice_speed"`    // 生成的语音速度
		VoicePitch   string `ini:"voice_pitch"`    // 生成的语音的语调
	} `ini:"ekho"` // ekho调度参数配置

	Logic struct {
		SoftwareWarehouse string `ini:"software_warehouse"` // 软件仓库位置
	} `ini:"logic"`
}
