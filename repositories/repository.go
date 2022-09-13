package repositories

import (
	"database/sql"
	_ "database/sql"
	"github.com/sirupsen/logrus"
	"video_storage/repositories/constants"
	"video_storage/sdk"
)

var (
	VideoRepository         *videoRepository
	UserRecordRepository    *userRecordRepository
	MemoryStorageRepository *memoryStorageRepository
	DictionaryRepository    *dictionaryRepository
)

// 初始化
func Init() {
	UserRecordRepository = &userRecordRepository{}
	_ = UserRecordRepository.repositoryTable(true)
	DictionaryRepository = &dictionaryRepository{}
	_ = DictionaryRepository.repositoryTable(true)
	VideoRepository = &videoRepository{}
	_ = VideoRepository.repositoryTable(true)
}

// 表检测
func tableRepository(tableName, createTableSQL string, needCreate bool) {
	args := []interface{}{"table", tableName}
	logrus.Info("创建", tableName, "表")
	sdk.SQLiteSDK.QueryOne(func(row sql.Row) {
		var rowCount int
		_ = row.Scan(&rowCount)
		if 0 == rowCount && needCreate {
			sdk.SQLiteSDK.Execute(nil, createTableSQL)
		}
	}, constants.RepositoryTable, args...)
}
