package routes

import (
	"video_storage/logic"
	"video_storage/model"
	"video_storage/tools"

	"github.com/kataras/iris/v12"
)

type VideoRoute struct {
	Ctx iris.Context
}

func (that *VideoRoute) GetPlayer() *model.ResponseBody {
	id, err := tools.FormValueInt64(that.Ctx, "id")
	if nil != err {
		return tools.Failed(400, err.Error())
	}
	if stream, err := logic.VideoLogic.PlayVideo(id); nil != err {
		return tools.Failed(400, err.Error())
	} else {
		return tools.Success(stream)
	}
}

// 创建视频
func (that *VideoRoute) Post() *model.ResponseBody {
	demandVideo := &model.DemandVideo{
		ID:       tools.FormValueInt64Default(that.Ctx, "id", 0),
		Describe: tools.FormValue(that.Ctx, "describe"),
		Duration: tools.FormValueFloat64Default(that.Ctx, "duration", 0),
		Height:   tools.FormValueIntDefault(that.Ctx, "height", 0),
		Size:     tools.FormValueInt64Default(that.Ctx, "size", 0),
		Title:    tools.FormValue(that.Ctx, "title"),
		Width:    tools.FormValueIntDefault(that.Ctx, "width", 0),
	}
	if err := logic.VideoLogic.NewVideo(demandVideo); nil != err {
		return tools.Failed(400, err.Error())
	}
	return tools.Success(nil)
}

// 上新
func (that *VideoRoute) PostFfmpeg() *model.ResponseBody {
	videoPath := tools.FormValue(that.Ctx, "path")
	video, err := logic.VideoLogic.FFmpeg(videoPath)
	if nil != err {
		return tools.Failed(400, err.Error())
	} else {
		return tools.Success(video)
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
