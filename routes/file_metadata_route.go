package routes

import (
	"video_storage/logic"
	"video_storage/model"
	"video_storage/tools"

	"github.com/kataras/iris/v12"
)

type FileMetadataRoute struct {
	Ctx iris.Context
}

func (that *FileMetadataRoute) GetAll() *model.ResponseBody {
	page := tools.FormValueIntDefault(that.Ctx, "page", 0)
	count := tools.FormValueIntDefault(that.Ctx, "count", 10)
	page -= 1
	if page < 0 {
		page = 0
	}
	return tools.Success(logic.FileMetadataLogic.FindAll(page, count))
}

func (that *FileMetadataRoute) GetSearch() *model.ResponseBody {
	page := tools.FormValueIntDefault(that.Ctx, "page", 0)
	count := tools.FormValueIntDefault(that.Ctx, "count", 10)
	page -= 1
	if page < 0 {
		page = 0
	}
	keyword := tools.FormValue(that.Ctx, "keyword")
	content, err := logic.FileMetadataLogic.FindByKeyword(keyword, page, count)
	if nil != err {
		return tools.Failed(400, err.Error())
	}

	return tools.Success(content)
}
