package repositories

import (
	"database/sql"
	"errors"
	"video_storage/model"
	"video_storage/repositories/constants"
	"video_storage/sdk"
	"video_storage/tools"
)

type videoRepository struct {
}

func (*videoRepository) repositoryTable(needCreate bool) error {
	tableRepository(constants.TableVideo, constants.CreateVideoTable, needCreate)
	return nil
}

func (*videoRepository) TotalCount() int {
	count := new(int)
	sdk.SQLiteSDK.QueryOne(func(row sql.Row) {
		_ = row.Scan(count)
	}, constants.VideoTotalCount, false)
	return *count
}

func (*videoRepository) GetList(begin, count int) []*model.DemandVideo {
	var videoList []*model.DemandVideo
	args := []interface{}{false, count, begin}
	sdk.SQLiteSDK.Query(func(rows sql.Rows) {
		defer rows.Close()
		columns, _ := rows.Columns()
		for rows.Next() {
			instance := new(model.DemandVideo)
			_ = rows.Scan(sdk.SQLiteSDK.ResultToModel(columns, instance)...)
			videoList = append(videoList, instance)
		}
	}, constants.VideoGetList, args...)
	return videoList
}

func (*videoRepository) FindByID(id int64) (*model.DemandVideo, error) {
	var err error
	demandVideo := &model.DemandVideo{}
	sdk.SQLiteSDK.Query(func(rows sql.Rows) {
		defer rows.Close()
		columns, _ := rows.Columns()
		if rows.Next() {
			_ = rows.Scan(sdk.SQLiteSDK.ResultToModel(columns, demandVideo)...)
		} else {
			err = errors.New("没有找到该视频")
		}
	}, constants.VideoFindByID, id)
	return demandVideo, err
}

func (*videoRepository) IsIncluded(videoPath string) (*model.DemandVideo, error) {
	var err error
	demandVideo := &model.DemandVideo{}
	sdk.SQLiteSDK.Query(func(rows sql.Rows) {
		defer rows.Close()
		columns, _ := rows.Columns()
		if rows.Next() {
			_ = rows.Scan(sdk.SQLiteSDK.ResultToModel(columns, demandVideo)...)
		} else {
			err = errors.New("视频未收录")
		}
	}, constants.VideoIsIncluded, videoPath)
	return demandVideo, err
}

func (*videoRepository) Save(demandVideo *model.DemandVideo) error {
	var err error
	var executeSQL string
	var args []interface{}
	if 0 == demandVideo.ID {
		executeSQL = constants.InsertDemandVideo
		demandVideo.CreateTime = tools.GetTime()
		demandVideo.ModifyTime = demandVideo.CreateTime
		args = []interface{}{demandVideo.Name, demandVideo.Size, demandVideo.Width, demandVideo.Height, demandVideo.Duration, demandVideo.Path, demandVideo.Describe, demandVideo.Title, demandVideo.FFmpegJSON, demandVideo.CreateTime, demandVideo.ModifyTime, false}
	} else {
		executeSQL = constants.UpdateDemandVideo
		demandVideo.ModifyTime = tools.GetTime()
		args = []interface{}{demandVideo.Name, demandVideo.Size, demandVideo.Width, demandVideo.Height, demandVideo.Duration, demandVideo.Path, demandVideo.Describe, demandVideo.Title, demandVideo.FFmpegJSON, demandVideo.ModifyTime, demandVideo.IsDel, demandVideo.ID}
	}
	sdk.SQLiteSDK.Execute(func(result sql.Result) {
		changeNumber, _ := result.RowsAffected()
		id, _ := result.LastInsertId()
		if 0 == changeNumber {
			err = errors.New("执行更新失败")
		}
		if 0 == demandVideo.ID && 0 == id {
			err = errors.New("插入失败")
		} else {
			demandVideo.ID = id
		}
	}, executeSQL, args...)
	return err
}