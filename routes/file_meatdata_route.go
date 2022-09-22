package routes

import (
	"video_storage/logic"
	"video_storage/model"
	"video_storage/tools"

	"github.com/kataras/iris/v12"
)

type FileMeatdataRoute struct {
	Ctx iris.Context
}

func (that *FileMeatdataRoute) GetAll() *model.ResponseBody {
	page := tools.FormValueIntDefault(that.Ctx, "page", 0)
	count := tools.FormValueIntDefault(that.Ctx, "count", 10)
	page -= 1
	if page < 0 {
		page = 0
	}
	return tools.Success(logic.FileMeatdataLogic.FindAll(page, count))
}
