package repositories

import (
	"database/sql"
	_ "database/sql"
	"video_storage/repositories/constants"
	"video_storage/sdk"
)

// 初始化
func Init() {
	UserRecordRepository = &userRecordRepository{}
	_ = UserRecordRepository.repositoryTable(true)
}

// 表检测
func tableRepository(tableName, createTableSQL string, needCreate bool) {
	args := []interface{}{"table", tableName}
	sdk.SQLiteSDK.QueryOne(func(row sql.Row) {
		var rowCount int
		_ = row.Scan(&rowCount)
		if 0 == rowCount && needCreate {
			sdk.SQLiteSDK.Execute(nil, createTableSQL)
		}
	}, constants.RepositoryTable, args...)
}