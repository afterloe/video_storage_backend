package logic

import (
	"video_storage/model"
	"video_storage/repositories"
	"video_storage/tools"
)

type dictionaryLogic struct {}

func (*dictionaryLogic) CreateDictionary(name, data string, groupID int64) error {
	if _, err := repositories.DictionaryRepository.FindDictionaryGroupByID(groupID); nil != err {
		return err
	}
	dictionary := &model.Dictionary{}
	dictionary.Name = name
	dictionary.Data = data
	dictionary.GroupID = groupID
	dictionary.IsDel = false
	dictionary.CreateTime = tools.GetTime()
	dictionary.ModifyTime = dictionary.CreateTime
	return repositories.DictionaryRepository.CreateDictionary(dictionary)
}

func (*dictionaryLogic) CreateGroup(name, groupType string) error {
	instance, err := repositories.DictionaryRepository.FindDictionaryGroupByName(name)
	if nil != err {
		return err
	}
	instance.IsDel = false
	instance.Name = name
	instance.GroupType = groupType
	instance.CreateTime = tools.GetTime()
	instance.ModifyTime = instance.CreateTime
	return repositories.DictionaryRepository.CreateDictionaryGroup(instance)
}