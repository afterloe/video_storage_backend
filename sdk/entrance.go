package sdk

import (
	"github.com/sirupsen/logrus"
	"sync"
)

var (
	SQLiteSDK sqliteSDK
)

func Init(sqliteFile string) {
	var once sync.Once
	once.Do(func() {
		logrus.Info("SDK初始化， 加载license文件")

		// sqlite sdk
		sqlSDK, _ := initSQLiteSDK(sqliteFile)

		// 赋值并初始化完毕
		SQLiteSDK = *sqlSDK
		logrus.Info("SDK初始化成功.")
	})
}

func CloseALLLink() error {
	SQLiteSDK.Disconnection()
	logrus.Info("数据库连接已关闭")
	return nil
}
