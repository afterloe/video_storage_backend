package logic

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"video_storage/config"
	"video_storage/logic/constants"
	"video_storage/model"
	"video_storage/repositories"
	"video_storage/tools"
)

type videoLogic struct{}

func (*videoLogic) PlayVideo(id int64) (interface{}, error) {
	// demandVideo, err := repositories.VideoRepository.FindByID(id)
	// if nil != err {
	// 	return nil, err
	// }
	// return fmt.Sprintf("/%s/%s", config.Instance.Logic.VideoPrefix, demandVideo.Name), err
	return nil, nil
}

func (*videoLogic) FetchFileRealPath(source *model.FileMetadata) (string, error) {
	videoPath := source.FullPath
	videoPath = regexp.QuoteMeta(videoPath)
	videoPath = strings.ReplaceAll(videoPath, " ", `\ `)
	_, err := os.Stat(videoPath)
	if os.IsExist(err) {
		return "", errors.New("源文件不存在，请检查源文件或对源数据进行调整和修改")
	}

	return videoPath, nil
}

func (that *videoLogic) SaveDescribe(info *model.VideoDescribe, source *model.FileMetadata) error {
	videoPath, err := that.FetchFileRealPath(source)
	if nil != err {
		return err
	}
	info.IsDel = false
	info.CreateTime = tools.GetTime()
	info.ModifyTime = info.CreateTime
	// 截取第10秒的视频截图
	receive := tools.ExecuteWithOutError(constants.FfmpegCatch, "10", videoPath, fmt.Sprintf("%s/%s", config.Instance.Logic.Screenshot, source.HexCode))
	if receive.HasError() {
		return receive.GetError()
	}
	return repositories.VideoRepository.Save(info)
}

func (that *videoLogic) Ffprobe(source *model.FileMetadata) (string, error) {
	videoPath, err := that.FetchFileRealPath(source)
	if nil != err {
		return "", err
	}
	receive := tools.Execute(constants.FfmpegJSON, videoPath)
	if receive.HasError() {
		return "", receive.GetError()
	}
	return receive.ToString(), err
}

func (*videoLogic) FindVideoByTarget(videoType string, page, count int) (*model.ListBody, error) {
	videoType = strings.Trim(videoType, " ")
	if videoType == "" {
		return nil, errors.New("搜索内容不能为空")
	}
	if isOK, _ := regexp.MatchString("^[\u4E00-\u9FA5A-Za-z0-9_]+$", videoType); !isOK {
		return nil, errors.New("输入的为非法字符")
	}
	dataList, total := repositories.VideoRepository.FindVideoListByTarget(videoType, count*page, count, false)
	return &model.ListBody{
		Total: total,
		Data:  dataList,
	}, nil
}
