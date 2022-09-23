package logic

import (
	"os"
	"regexp"
	"strings"
	"video_storage/model"
	"video_storage/repositories"
	"video_storage/tools"

	"github.com/sirupsen/logrus"
)

type videoLogic struct {
}

func (*videoLogic) PlayVideo(id int64) (interface{}, error) {
	// demandVideo, err := repositories.VideoRepository.FindByID(id)
	// if nil != err {
	// 	return nil, err
	// }
	// return fmt.Sprintf("/%s/%s", config.Instance.Logic.VideoPrefix, demandVideo.Name), err
	return nil, nil
}

func (*videoLogic) SaveDescribe(info *model.VideoDescribe) error {
	info.IsDel = false
	info.CreateTime = tools.GetTime()
	info.ModifyTime = info.CreateTime
	return repositories.VideoRepository.Save(info)
}

func (*videoLogic) Ffprobe(videoPath string) (string, error) {
	var err error
	_, err = os.Stat(videoPath)
	if nil != err {
		logrus.Error(err)
		return "", err
	}
	videoPath = regexp.QuoteMeta(videoPath)
	videoPath = strings.ReplaceAll(videoPath, " ", `\ `)
	receive := tools.Execute("ffprobe -v quiet -print_format json -show_format -show_streams %s", videoPath)
	if receive.HasError() {
		return "", receive.GetError()
	}
	return receive.ToString(), err
}

func (*videoLogic) FindVideoByTarget(videoType string, page, count int) *model.ListBody {
	// dataList := repositories.VideoRepository.GetList(count*page, count)
	// totalNumber := repositories.VideoRepository.TotalCount()
	return &model.ListBody{
		Total: 0,
		Data:  nil,
	}
}
