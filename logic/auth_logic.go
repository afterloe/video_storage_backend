package logic

import (
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
	"video_storage/model"
	"video_storage/repositories"
)

// 权限拦截 中间件
func AuthLogic(ctx iris.Context) {
	token := ctx.FormValue("token")
	if 0 == len(token) {
		token = ctx.GetHeader("token")
	}
	if pointer, err := repositories.MemoryStorageRepository.Get("user", token); nil != err {
		pleaseSignIn(ctx)
		ctx.StopExecution()
	} else {
		user := pointer.(*model.User)
		// 打印访问日志
		logrus.Infof("uid: %d use %s accept %s", user.ID, ctx.Method(), ctx.RequestPath(true))

		// TODO 权限拦截
		ctx.Values().Set("token", token)
		ctx.Values().Set("uid", user.ID)
	}
	ctx.Next()
}


// pleaseSignIn 未登录返回提示
func pleaseSignIn(ctx iris.Context) {
	_, _ = ctx.JSON(&model.ResponseBody{
		Code:    401,
		Message: "请登录",
	})
	ctx.StopExecution()
}

// 获取登录用户信息
func WhoYouAre(ctx iris.Context) *model.User {
	uid := ctx.Values().Get("uid")
	if 0 == uid {
		pleaseSignIn(ctx)
		return nil
	}
	user := repositories.UserRecordRepository.FindByID(uid)
	user.Passwd = ""
	return user
}