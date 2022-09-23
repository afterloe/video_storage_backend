package repositories

import (
	"database/sql"
	"errors"
	"video_storage/model"
	"video_storage/repositories/constants"
	"video_storage/sdk"
)

type objectRepository struct {
}

func (*objectRepository) repositoryTable(needCreate bool) error {
	tableRepository(constants.TableObject, constants.CreateObjectTable, needCreate)
	return nil
}

func (*objectRepository) FindObjectByMetadataID(id int64) *model.Object {
	args := []interface{}{id}
	instance := &model.Object{}
	sdk.SQLiteSDK.QueryOne(func(r sql.Row) {
		_ = r.Scan(sdk.SQLiteSDK.ResultToModelBySQL(constants.FindObjectByMetadataID, instance)...)
		if instance.ID == 0 {
			instance = nil
		}
	}, constants.FindObjectByMetadataID, args...)

	return instance
}

func (*objectRepository) Save(object *model.Object) error {
	var (
		err        error
		executeSQL string
		args       []interface{}
	)
	if object.ID == 0 {
		executeSQL = constants.InsertObject
		args = []interface{}{object.VirtualPath, object.MetadataID, object.CreateTime, object.IsDel, object.ModifyTime}
	} else {

	}

	sdk.SQLiteSDK.Execute(func(result sql.Result) {
		changeNumber, _ := result.RowsAffected()
		id, _ := result.LastInsertId()
		if changeNumber == 0 {
			err = errors.New("执行更新失败")
		}
		if object.ID == 0 && id == 0 {
			err = errors.New("插入失败")
		} else {
			object.ID = id
		}
	}, executeSQL, args...)

	return err
}
