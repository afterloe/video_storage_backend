package routes

import (
	"github.com/kataras/iris/v12"
	"video_storage/logic"
	"video_storage/model"
	"video_storage/tools"
)

type UserRoute struct {
	Ctx iris.Context
}

func (that *UserRoute) GetMe() *model.ResponseBody {
	user := logic.WhoYouAre(that.Ctx)
	return tools.Success(user)
}

func (that *UserRoute) GetPing() *model.ResponseBody {
	return tools.Success(nil)
}

func (that *UserRoute) DeleteCancellation() *model.ResponseBody {
	logic.UserLogic.Cancellation(that.Ctx.Values().Get("token").(string))
	return tools.Success(nil)
}
