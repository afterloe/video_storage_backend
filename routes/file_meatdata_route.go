package routes

import (
	"video_storage/model"
	"video_storage/repositories"
	"video_storage/tools"

	"github.com/kataras/iris/v12"
)

type FileMeatdataRoute struct {
	Ctx iris.Context
}

func (that *FileMeatdataRoute) GetAll() *model.ResponseBody {
	fileList := repositories.FileMeatdataRepository.FindAll(0, 10)
	return tools.Success(fileList)
}

// func (that *FileMeatdataRoute) PostAll() *model.ResponseBody {
// 	begin, _ := tools.FormValueInt(that.Ctx, "begin")
// 	count, _ := tools.FormValueInt(that.Ctx, "count")
// 	fileList := repositories.FileMeatdataRepository.FindAll(begin, count)
// 	r := &model.LoginBody{}
// 	r.Token = "info"
// 	r.User = fileList
// 	return tools.Success(r)
// }
