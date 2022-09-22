package logic

import (
	"video_storage/model"
	"video_storage/repositories"
	"video_storage/tools"
)

type dictionaryLogic struct{}

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

func (*dictionaryLogic) DeleteDictionary(dictionaryID int) error {
	_, err := repositories.DictionaryRepository.FindDictionaryByID(dictionaryID)
	if nil != err {
		return err
	}
	err = repositories.DictionaryRepository.DeleteDictionary(dictionaryID)
	return err
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

func (*dictionaryLogic) GetDictionaryGroupList() []*model.DictionaryGroup {
	list := repositories.DictionaryRepository.FindAllDictionaryGroup()
	repositories.DictionaryRepository.FindAllDictionary(list)
	return list
}

func (*dictionaryLogic) GetDictionaryGroup(dictionaryType string) []*model.Dictionary {
	group := repositories.DictionaryRepository.FindDictionaryGroupByGroupType(dictionaryType)
	if 0 == group.ID {
		return nil
	}
	repositories.DictionaryRepository.FindAllDictionary([]*model.DictionaryGroup{group})
	return group.Values
}
