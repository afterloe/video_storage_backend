package logic

import (
	"video_storage/repositories"
	"video_storage/tools"
)

type dictionaryLogic struct {
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