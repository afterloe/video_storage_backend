package repositories

import (
	"database/sql"
	"video_storage/model"
	"video_storage/repositories/constants"
	"video_storage/sdk"
)

type fileMetadataRepository struct {
}

func (*fileMetadataRepository) repositoryTable(needCreate bool) error {
	tableRepository(constants.TableFileMetadata, constants.CreateFileMetaDataTable, needCreate)
	return nil
}

func (*fileMetadataRepository) TotalCount() int {
	count := new(int)
	sdk.SQLiteSDK.QueryOne(func(row sql.Row) {
		_ = row.Scan(count)
	}, constants.FileMetadataCount, false)
	return *count
}

func (*fileMetadataRepository) FindAll(start, count int) []*model.FileMetadata {
	var metadataList []*model.FileMetadata
	args := []interface{}{count, start}
	sdk.SQLiteSDK.Query(func(r sql.Rows) {
		defer r.Close()
		columns, _ := r.Columns()
		for r.Next() {
			i := new(model.FileMetadata)
			_ = r.Scan(sdk.SQLiteSDK.ResultToModel(columns, i)...)
			metadataList = append(metadataList, i)
		}
	}, constants.FindAllFileMetaData, args...)

	return metadataList
}

func (*fileMetadataRepository) FindByKeyword(keyword string, start, count int, isDel bool) ([]*model.FileMetadata, int) {
	var metadataList []*model.FileMetadata
	var total int
	args := []interface{}{keyword, isDel, count, start}
	sdk.SQLiteSDK.Query(func(rows sql.Rows) {
		defer rows.Close()
		columns, _ := rows.Columns()
		for rows.Next() {
			i := new(model.FileMetadata)
			_ = rows.Scan(sdk.SQLiteSDK.ResultToModel(columns, i)...)
			metadataList = append(metadataList, i)
		}
	}, constants.FindFileMetadataByKeyword, args...)
	sdk.SQLiteSDK.QueryOne(func(row sql.Row) {
		_ = row.Scan(&total)
	}, constants.FindFileMetadataByKeywordCount, args...)
	return metadataList, total
}
