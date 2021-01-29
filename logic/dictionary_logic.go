package logic

import "video_storage/repositories"

type dictionaryLogic struct {
}

func (*dictionaryLogic) CreateGroup(name, groupType string) error {
	instance, err := repositories.DictionaryRepository.FindDictionaryGroupByName(name)
	if nil != err {
		return err
	}
	// TODO
	return nil
}