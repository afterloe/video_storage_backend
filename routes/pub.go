package routes

import (
	"github.com/kataras/iris/v12"
	"video_storage/config"
	"video_storage/model"
	"video_storage/repositories"
	"video_storage/tools"
)

type PubRoute struct {
	Ctx iris.Context
}

// 获取服务器目前版本
func (*PubRoute) GetVersion() *model.ResponseBody {
	return tools.Success(config.Instance.Common.Version)
}

// 注册
func (that *PubRoute) PostSignUp() *model.ResponseBody {
	email := tools.FormValue(that.Ctx, "email")
	password := tools.FormValue(that.Ctx, "password")
	if "" == email {
		return tools.Failed(400, "email 不能为空")
	}
	if user, err := repositories.UserRecordRepository.SignUp(email, password); nil != err {
		return tools.Failed(400, err.Error())
	} else {
		return tools.Success(user)
	}
}

// 登陆
func (that *PubRoute) PutSignIn() *model.ResponseBody {
	email := tools.FormValue(that.Ctx, "email")
	password := tools.FormValue(that.Ctx, "password")
	if "" == email {
		return tools.Failed(400, "email 不能为空")
	}
	if user, err := repositories.UserRecordRepository.SignIn(email, password); nil != err {
		return tools.Failed(400, err.Error())
	} else {
		token := repositories.MemoryStorageRepository.Set("user", user)
		return tools.Success(token)
	}
}
