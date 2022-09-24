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
	tableRepository(constants.ViewVideoDescribePackage, constants.CreateDescribePackageView, needCreate)
	return nil
}

func (*videoRepository) FindVideoListByTarget(videoType string, start, count int, isDel bool) ([]*model.VideoDescribePackage, int) {
	var (
		videoList []*model.VideoDescribePackage
		total     int
	)
	args := []interface{}{count, start}
	sdk.SQLiteSDK.Query(func(r sql.Rows) {
		defer r.Close()
		columns, _ := r.Columns()
		for r.Next() {
			i := new(model.VideoDescribePackage)
			_ = r.Scan(sdk.SQLiteSDK.ResultToModel(columns, i)...)
			videoList = append(videoList, i)
		}
	}, constants.FindVideoDescribePackageByTag, args...)

	sdk.SQLiteSDK.QueryOne(func(row sql.Row) {
		_ = row.Scan(&total)
	}, constants.FindVideoDescribePackageByTagCount, args...)

	return videoList, total
}

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
		// TODO
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
