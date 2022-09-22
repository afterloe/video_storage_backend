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

func (*dictionaryLogic) UpdateDictionary(id int64, name, data string) (*model.Dictionary, error) {
	instance, err := repositories.DictionaryRepository.FindDictionaryByID(id)
	if nil != err {
		return nil, err
	}
	instance.Name = name
	instance.Data = data
	instance.ModifyTime = tools.GetTime()
	err = repositories.DictionaryRepository.ModifyDictionary(instance)
	return instance, err
}

func (*dictionaryLogic) DeleteDictionary(dictionaryID int64) error {
	_, err := repositories.DictionaryRepository.FindDictionaryByID(dictionaryID)
	if nil != err {
		return err
	}
	err = repositories.DictionaryRepository.DeleteDictionary(dictionaryID)
	return err
}

func (*dictionaryLogic) CreateGroup(name, groupType string) (*model.DictionaryGroup, error) {
	instance, err := repositories.DictionaryRepository.FindDictionaryGroupByName(name)
	if nil != err {
		return nil, err
	}
	instance.IsDel = false
	instance.Name = name
	instance.GroupType = groupType
	instance.CreateTime = tools.GetTime()
	instance.ModifyTime = instance.CreateTime
	err = repositories.DictionaryRepository.CreateDictionaryGroup(instance)
	return instance, err
}

func (*dictionaryLogic) UpdateGroup(id int64, name, groupType string) (*model.DictionaryGroup, error) {
	instance, err := repositories.DictionaryRepository.FindDictionaryGroupByID(id)
	if nil != err {
		return nil, err
	}
	instance.Name = name
	instance.GroupType = groupType
	instance.ModifyTime = tools.GetTime()
	err = repositories.DictionaryRepository.ModifyDictionaryGroup(instance)
	return instance, err
}

func (*dictionaryLogic) DeleteGroup(groupID int64) error {
	_, err := repositories.DictionaryRepository.FindDictionaryGroupByID(groupID)
	if nil != err {
		return err
	}
	err = repositories.DictionaryRepository.DeleteDictionaryGroup(groupID)
	return err
}

func (*dictionaryLogic) GetDictionaryGroupList() []*model.DictionaryGroup {
	list := repositories.DictionaryRepository.FindAllDictionaryGroup()
	repositories.DictionaryRepository.FindAllDictionary(list)
	return list
}

func (*dictionaryLogic) GetDictionaryGroup(dictionaryType string) []*model.Dictionary {
	group := repositories.DictionaryRepository.FindDictionaryGroupByGroupType(dictionaryType)
	if group.ID == 0 {
		return nil
	}
	repositories.DictionaryRepository.FindAllDictionary([]*model.DictionaryGroup{group})
	return group.Values
}
