package logic

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strings"
	"video_storage/model"
	"video_storage/repositories"
	"video_storage/tools"
)

type videoLogic struct {
}

func (*videoLogic) NewVideo(instance *model.DemandVideo) error {
	var err error
	demandVideo, err := repositories.VideoRepository.FindByID(instance.ID)
	if nil != err {
		return err
	}
	demandVideo.ModifyTime = tools.GetTime()
	demandVideo.Describe = instance.Describe
	demandVideo.Duration = instance.Duration
	demandVideo.Height = instance.Height
	demandVideo.Width = instance.Width
	demandVideo.Size = instance.Size
	demandVideo.Title = instance.Title
	err = repositories.VideoRepository.Save(demandVideo)
	return err
}

func (*videoLogic) FFmpeg(videoPath string) (*model.DemandVideo, error) {
	var err error
	_, err = os.Stat(videoPath)
	if nil != err {
		logrus.Error(err)
		return nil, err
	}
	demandVideo, err := repositories.VideoRepository.IsIncluded(videoPath)
	if nil == err {
		return demandVideo, nil
	}
	ffmpegJSON, err := tools.Execute(fmt.Sprintf("ffprobe -select_streams v \\\n-show_entries format=duration,size,bit_rate,filename \\\n-show_streams \\\n-v quiet \\\n-of csv=\"p=0\" \\\n-of json \\\n-i %s", videoPath))
	if nil != err {
		err = errors.New("获取ffmpeg信息失败")
		return nil, err
	}
	demandVideo.FFmpegJSON = ffmpegJSON
	demandVideo.Path = videoPath
	demandVideo.Name = tools.GeneratorUUID()
	err = repositories.VideoRepository.Save(demandVideo)
	return demandVideo, err
}

func (*videoLogic) FindVideoByTarget(videoType string, page, count int) map[string]interface{} {
	// TODO
	return map[string]interface{}{}
}

func (*videoLogic) ScanVideo(path string) ([]*model.ScanFile, error) {
	info, err := os.Stat(path)
	if nil != err {
		if os.IsNotExist(err) {
			return nil, errors.New("文件目录不存在")
		}
	}
	var videoList []*model.ScanFile
	scanFile(path, info, &videoList)
	return videoList, nil
}

func scanFile(path string, info os.FileInfo, list *[]*model.ScanFile) {
	if info.IsDir() {
		childItem, _ := ioutil.ReadDir(path)
		for _, i := range childItem {
			scanFile(path+"/"+i.Name(), i, list)
		}
	} else {
		f := &model.ScanFile{
			Name:       info.Name(),
			Path:       path,
			Mode:       info.Mode(),
			Size:       info.Size(),
			ModifyTime: tools.FormatTime(info.ModTime()),
		}
		if 0 == strings.Index(f.Name, ".") {
			return
		}
		*list = append(*list, f)
	}
}
