package routes

import (
	"github.com/kataras/iris/v12"
	"video_storage/logic"
	"video_storage/model"
	"video_storage/tools"
)

type VideoRoute struct {
	Ctx iris.Context
}

// 指定扫描
func (that *VideoRoute) PostScan() *model.ResponseBody {
	inputPath := tools.FormValue(that.Ctx, "path")
	if videoList, err := logic.VideoLogic.ScanVideo(inputPath); nil != err {
		return tools.Failed(400, err.Error())
	} else {
		return tools.Success(videoList)
	}
}

// 视频目录
func (that *VideoRoute) GetList() *model.ResponseBody {
	videoType := tools.FormValueDefault(that.Ctx, "type", "hot")
	page := tools.FormValueIntDefault(that.Ctx, "page", 0)
	count := tools.FormValueIntDefault(that.Ctx, "count", 10)
	page -= 1
	if page < 0 {
		page = 0
	}
	return tools.Success(logic.VideoLogic.FindVideoByTarget(videoType, page, count))
}
