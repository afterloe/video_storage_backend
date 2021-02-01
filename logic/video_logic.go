package logic

import (
	"errors"
	"io/ioutil"
	"os"
	"video_storage/model"
	"video_storage/tools"
)

type videoLogic struct {
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
		*list = append(*list, f)
	}
}
