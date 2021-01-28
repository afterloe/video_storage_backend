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
