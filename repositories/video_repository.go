package repositories

import (
	"database/sql"
	"errors"
	"video_storage/model"
	"video_storage/repositories/constants"
	"video_storage/sdk"
)

type videoRepository struct {
}

func (*videoRepository) repositoryTable(needCreate bool) error {
	tableRepository(constants.TableVideo, constants.CreateVideoTable, needCreate)
	return nil
}

// func (*videoRepository) TotalCount() int {
// 	count := new(int)
// 	sdk.SQLiteSDK.QueryOne(func(row sql.Row) {
// 		_ = row.Scan(count)
// 	}, constants.VideoTotalCount, false)
// 	return *count
// }

// func (*videoRepository) GetList(begin, count int) []*model.DemandVideo {
// 	var videoList []*model.DemandVideo
// 	args := []interface{}{false, count, begin}
// 	sdk.SQLiteSDK.Query(func(rows sql.Rows) {
// 		defer rows.Close()
// 		columns, _ := rows.Columns()
// 		for rows.Next() {
// 			instance := new(model.DemandVideo)
// 			_ = rows.Scan(sdk.SQLiteSDK.ResultToModel(columns, instance)...)
// 			videoList = append(videoList, instance)
// 		}
// 	}, constants.VideoGetList, args...)
// 	return videoList
// }

// func (*videoRepository) FindByID(id int64) (*model.DemandVideo, error) {
// 	var err error
// 	demandVideo := &model.DemandVideo{}
// 	sdk.SQLiteSDK.Query(func(rows sql.Rows) {
// 		defer rows.Close()
// 		columns, _ := rows.Columns()
// 		if rows.Next() {
// 			_ = rows.Scan(sdk.SQLiteSDK.ResultToModel(columns, demandVideo)...)
// 		} else {
// 			err = errors.New("没有找到该视频")
// 		}
// 	}, constants.VideoFindByID, id)
// 	return demandVideo, err
// }

// func (*videoRepository) IsIncluded(videoPath string) (*model.DemandVideo, error) {
// 	var err error
// 	demandVideo := &model.DemandVideo{}
// 	sdk.SQLiteSDK.Query(func(rows sql.Rows) {
// 		defer rows.Close()
// 		columns, _ := rows.Columns()
// 		if rows.Next() {
// 			_ = rows.Scan(sdk.SQLiteSDK.ResultToModel(columns, demandVideo)...)
// 		} else {
// 			err = errors.New("视频未收录")
// 		}
// 	}, constants.VideoIsIncluded, videoPath)
// 	return demandVideo, err
// }

func (*videoRepository) Save(describe *model.VideoDescribe) error {
	var (
		err        error
		executeSQL string
		args       []interface{}
	)
	if describe.ID == 0 {
		executeSQL = constants.InsertVideoDescribe
		args = []interface{}{describe.MetadataID, describe.Width, describe.Height, describe.Duration, describe.CodecName, describe.DisplayAspectRatio, describe.CodecLongName, describe.CreateTime, describe.ModifyTime, describe.IsDel}
	} else {

	}
	sdk.SQLiteSDK.Execute(func(result sql.Result) {
		changeNumber, _ := result.RowsAffected()
		id, _ := result.LastInsertId()
		if changeNumber == 0 {
			err = errors.New("执行更新失败")
		}
		if describe.ID == 0 && id == 0 {
			err = errors.New("插入失败")
		} else {
			describe.ID = id
		}
	}, executeSQL, args...)
	return err
}
