package logic

import (
	"errors"
	"regexp"
	"strings"
	"video_storage/config"
	"video_storage/model"
	"video_storage/repositories"
	"video_storage/tools"
)

type objectLogic struct {
}

func (*objectLogic) FindByID(id int64) (*model.Object, error) {
	object := repositories.ObjectRepository.FindObjectByMetadataID(id)
	if object == nil {
		return object, errors.New("未检索到该id对应的对象信息")
	}

	return object, nil
}

func (*objectLogic) SaveObject(source *model.FileMetadata) error {
	object := &model.Object{}
	object.IsDel = false
	object.CreateTime = tools.GetTime()
	object.ModifyTime = object.CreateTime
	object.MetadataID = source.ID
	object.VirtualPath = source.HexCode
	fullPath := regexp.QuoteMeta(source.FullPath)
	fullPath = strings.ReplaceAll(fullPath, " ", `\ `)
	receive := tools.Execute("ln -s %s %s/%s", fullPath, config.Instance.Logic.VideoStorage, object.VirtualPath)
	if receive.HasError() {
		return receive.GetError()
	}
	err := repositories.ObjectRepository.Save(object)
	if nil != err {
		source.IsLink = true
		err = repositories.FileMetadataRepository.Save(source)
	}

	return err
}
