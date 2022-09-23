package logic

import (
	"errors"
	"regexp"
	"strings"
	"video_storage/model"
	"video_storage/repositories"
)

type fileMetadataLogic struct{}

func (*fileMetadataLogic) FindAll(page, count int) *model.ListBody {
	dataList := repositories.FileMetadataRepository.FindAll(page*count, count)
	totalNumber := repositories.FileMetadataRepository.TotalCount()
	body := &model.ListBody{
		Data:  dataList,
		Total: totalNumber,
	}

	return body
}

func (*fileMetadataLogic) FindByKeyword(keyword string, page, count int) (*model.ListBody, error) {
	keyword = strings.Trim(keyword, " ")
	if keyword == "" {
		return nil, errors.New("搜索内容不能为空")
	}
	if isOK, _ := regexp.MatchString("^[\u4E00-\u9FA5A-Za-z0-9_]+$", keyword); !isOK {
		return nil, errors.New("输入的为非法字符")
	}
	dataList, total := repositories.FileMetadataRepository.FindByKeyword("%"+keyword+"%", page*count, count, false)
	body := &model.ListBody{
		Data:  dataList,
		Total: total,
	}

	return body, nil
}
