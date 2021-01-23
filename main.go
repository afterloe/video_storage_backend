package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"video_storage/config"
	"video_storage/tools"
)

const appVersion = "0.0.1"

var cfgFilePointer = flag.String("config", "./video_storage.ini.example", "配置文件路径")

func init() {
	startUpStr := fmt.Sprintf(`
      .__    .___                       __                                      
___  _|__| __| _/____  ____     _______/  |_  ________________     ____   ____  
\  \/ /  |/ __ |/ __ \/  _ \   /  ___/\   __\/  _ \_  __ \__  \   / ___\_/ __ \ 
 \   /|  / /_/ \  ___(  <_> )  \___ \  |  | (  <_> )  | \// __ \_/ /_/  >  ___/ 
  \_/ |__\____ |\___  >____/  /____  > |__|  \____/|__|  (____  /\___  / \___  >
              \/    \/             \/                         \//_____/      \/ 

version:\t%s
author:\tafterloe(605728727@qq.com)
`, appVersion)

	fmt.Println(startUpStr)
	flag.Parse()

	config.ReadInICfgFile(*cfgFilePointer)
	if stream := tools.GeneratorLogStream(config.Instance.Common.LogPath, "bootstrap"); nil != stream {
		if "debug" == config.Instance.Common.LogLevel {
			logrus.SetLevel(logrus.DebugLevel)
		} else {
			logrus.SetLevel(logrus.InfoLevel)
		}
	} else {
		logrus.SetOutput(os.Stdout)
		logrus.Error("日志写入失败， 输出到默认输出源")
	}
}

func main() {
	// 启动application

}
