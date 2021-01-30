package repositories

import (
	"database/sql"
	"errors"
	"video_storage/model"
	"video_storage/repositories/constants"
	"video_storage/sdk"
)

var DictionaryRepository *dictionaryRepository

type dictionaryRepository struct {
}

func (*dictionaryRepository) repositoryTable(needCreate bool) error {
	tableRepository(constants.TableDictionaryGroup, constants.CreateDictionaryGroupTable, needCreate)
	tableRepository(constants.TableDictionary, constants.CreateDictionaryTable, needCreate)
	return nil
}

func (*dictionaryRepository) FindDictionaryGroupByName(name string) (*model.DictionaryGroup, error) {
	var (
		instance model.DictionaryGroup
		err      error
	)
	args := []interface{}{name}
	sdk.SQLiteSDK.QueryOne(func(row sql.Row) {
		_ = row.Scan(&instance.ID, &instance.Name, &instance.GroupType)
		if 0 != instance.ID {
			err = errors.New("标签组已经被创建了")
		}
	}, constants.FindDictionaryGroupByName, args...)

	return &instance, err
}

func (*dictionaryRepository) CreateDictionaryGroup(instance *model.DictionaryGroup) error {
	var err error
	args := []interface{}{instance.Name, instance.GroupType, instance.CreateTime, instance.ModifyTime, instance.IsDel}
	sdk.SQLiteSDK.Execute(func(result sql.Result) {
		id, _ := result.LastInsertId()
		instance.ID = id
		if 0 == instance.ID {
			err = errors.New("插入失败")
		}
	}, constants.CreateDictionaryGroup, args...)

	return err
}

func (*dictionaryRepository) FindDictionaryGroupByID(id int64) (*model.DictionaryGroup, error) {
	var err error
	args := []interface{}{id}
	instance := &model.DictionaryGroup{}
	sdk.SQLiteSDK.QueryOne(func(row sql.Row) {
		_ = row.Scan(sdk.SQLiteSDK.ResultToModel([]string{"id", "name", "group_type", "create_time", "modify_time", "is_del"}, instance)...)
		if 0 == instance.ID {
			err = errors.New("group 不存在")
		}
	}, constants.FindDictionaryGroupByID, args...)
	return instance, err
}

func (*dictionaryRepository) CreateDictionary(instance *model.Dictionary) error {
	var err error
	args := []interface{}{instance.Name, instance.Data, instance.GroupID, instance.CreateTime, instance.ModifyTime, instance.IsDel}
	sdk.SQLiteSDK.Execute(func(result sql.Result) {
		id, _ := result.LastInsertId()
		instance.ID = id
		if 0 == instance.ID {
			err = errors.New("插入失败")
		}
	}, constants.CreateDictionary, args...)
	return err
}

func (*dictionaryRepository) FindAllDictionaryGroup() []*model.DictionaryGroup {
	var list []*model.DictionaryGroup
	sdk.SQLiteSDK.Query(func(rows sql.Rows) {
		defer rows.Close()
		columns, _ := rows.Columns()
		for rows.Next() {
			instance := new(model.DictionaryGroup)
			_ = rows.Scan(sdk.SQLiteSDK.ResultToModel(columns, instance)...)
			list = append(list, instance)
		}
	}, constants.FindAllDictionaryGroup, false)
	return list
}

func (*dictionaryRepository) FindAllDictionary(list []*model.DictionaryGroup) {
	for _, g := range list {
		var list []*model.Dictionary
		sdk.SQLiteSDK.Query(func(rows sql.Rows) {
			defer rows.Close()
			columns, _ := rows.Columns()
			for rows.Next() {
				instance := new(model.Dictionary)
				_ = rows.Scan(sdk.SQLiteSDK.ResultToModel(columns, instance)...)
				list = append(list, instance)
			}
		}, constants.FindAllDictionary, g.ID)
		g.Values = list
	}
}