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

func (*fileMetadataRepository) FindAll(begin, count int) []*model.FileMetadata {
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
