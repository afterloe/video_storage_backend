package sdk

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
)

var SQLiteSDK sqliteSDK

func Init(sqliteFilePath string) {
	var once sync.Once
	once.Do(func() {
		logrus.Info("SDK初始化")

		// sqlite sdk
		sqlSDK, _ := initSQLiteSDK(fmt.Sprintf("%s/videoStorage.db", sqliteFilePath))

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
