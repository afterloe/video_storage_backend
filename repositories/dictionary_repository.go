package repositories

import (
	"video_storage/model"
	"video_storage/repositories/constants"
)

var DictionaryRepository *dictionaryRepository

type dictionaryRepository struct {

}

func (*dictionaryRepository) repositoryTable(needCreate bool) error {
	tableRepository(constants.TableDictionary, constants.CreateDictionaryTable, needCreate)
	return nil
}

func (*dictionaryRepository) FindDictionaryGroupByName(name string) (*model.DictionaryGroup, error) {
	var (
		instance model.DictionaryGroup
		err error
	)
	// TODO
	return &instance, err
}