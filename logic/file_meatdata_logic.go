package logic

import (
	"video_storage/model"
	"video_storage/repositories"
)

type fileMeatdataLogic struct{}

func (*fileMeatdataLogic) FindAll(page, count int) *model.ListBody {
	dataList := repositories.FileMeatdataRepository.FindAll(page*count, count)
	totalNumber := repositories.FileMeatdataRepository.TotalCount()
	body := &model.ListBody{
		Data:  dataList,
		Total: totalNumber,
	}

	return body
}
