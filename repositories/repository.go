package repositories

import (
	"database/sql"
	"video_storage/repositories/constants"
	"video_storage/sdk"

	"github.com/sirupsen/logrus"
)

var (
	VideoRepository         *videoRepository
	UserRecordRepository    *userRecordRepository
	MemoryStorageRepository *memoryStorageRepository
	DictionaryRepository    *dictionaryRepository
	FileMetadataRepository  *fileMetadataRepository
	ObjectRepository        *objectRepository
)

// 初始化
func Init() {
	MemoryStorageRepository = &memoryStorageRepository{}
	UserRecordRepository = &userRecordRepository{}
	_ = UserRecordRepository.repositoryTable(true)
	DictionaryRepository = &dictionaryRepository{}
	_ = DictionaryRepository.repositoryTable(true)
	VideoRepository = &videoRepository{}
	_ = VideoRepository.repositoryTable(true)
	FileMetadataRepository = &fileMetadataRepository{}
	_ = FileMetadataRepository.repositoryTable(true)
	ObjectRepository = &objectRepository{}
	_ = ObjectRepository.repositoryTable(true)
}

// 表检测
func tableRepository(tableName, createTableSQL string, needCreate bool) {
	args := []interface{}{"table", tableName}
	logrus.Info("创建", tableName, "表")
	sdk.SQLiteSDK.QueryOne(func(row sql.Row) {
		var rowCount int
		_ = row.Scan(&rowCount)
		if rowCount == 0 && needCreate {
			sdk.SQLiteSDK.Execute(nil, createTableSQL)
		}
	}, constants.RepositoryTable, args...)
}
