package repositories

import (
	"database/sql"
	"video_storage/model"
	"video_storage/repositories/constants"
	"video_storage/sdk"
)

type fileMeatdataRepository struct {
}

func (*fileMeatdataRepository) repositoryTable(needCreate bool) error {
	tableRepository(constants.TableFileMetadata, constants.CreateFileMetaDataTable, needCreate)
	return nil
}

func (*fileMeatdataRepository) TotalCount() int {
	count := new(int)
	sdk.SQLiteSDK.QueryOne(func(row sql.Row) {
		_ = row.Scan(count)
	}, constants.FileMeatdataCount, false)
	return *count
}

func (*fileMeatdataRepository) FindAll(begin, count int) []*model.FileMetadata {
	var metadataList []*model.FileMetadata
	args := []interface{}{count, begin}
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
