package logic

import (
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
	"video_storage/model"
	"video_storage/repositories"
	"video_storage/tools"
)

// 权限拦截 中间件
func AuthLogic(ctx iris.Context) {
	token := ctx.FormValue("token")
	if 0 == len(token) {
		token = ctx.GetHeader("token")
	}
	if uid, err := repositories.MemoryStorageRepository.Get("user", token); nil != err {
		pleaseSignIn(ctx)
		ctx.StopExecution()
	} else {
		// 打印访问日志
		logrus.Infof("uid: %s use %s accept %s", uid, ctx.Method(), ctx.RequestPath(true))

		// TODO 权限拦截
		ctx.Values().Set("token", token)
		ctx.Values().Set("uid", uid)
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
	uid := ctx.Values().Get("uid").(string)
	if err := tools.CheckStr(uid); nil != err {
		pleaseSignIn(ctx)
		return nil
	}
	return repositories.UserRecordRepository.FindByID(uid)
}