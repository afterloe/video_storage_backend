package routes

import (
	"video_storage/config"
	"video_storage/logic"
	"video_storage/model"
	"video_storage/repositories"
	"video_storage/tools"

	"github.com/kataras/iris/v12"
)

type PubRoute struct {
	Ctx iris.Context
}

// 获取服务器目前版本
func (*PubRoute) GetVersion() *model.ResponseBody {
	return tools.Success(config.Instance.Common.Version)
}

// 注册
func (that *PubRoute) PutSignup() *model.ResponseBody {
	email := tools.FormValue(that.Ctx, "email")
	password := tools.FormValue(that.Ctx, "password")
	if email == "" {
		return tools.Failed(400, "email 不能为空")
	}
	if user, err := logic.UserLogic.SignUp(email, password); nil != err {
		return tools.Failed(400, err.Error())
	} else {
		return tools.Success(user)
	}
}

func (that *PubRoute) PostSignin1() *model.ResponseBody {
	email := tools.FormValue(that.Ctx, "email")
	password := tools.FormValue(that.Ctx, "password")
	if email == "" {
		return tools.Failed(400, "email 不能为空")
	}
	return tools.Success(map[string]interface{}{"email": email, "password": password})
}

// 登陆
func (that *PubRoute) PostSignin() *model.ResponseBody {
	email := tools.FormValue(that.Ctx, "email")
	password := tools.FormValue(that.Ctx, "password")
	if email == "" {
		return tools.Failed(400, "email 不能为空")
	}
	user, err := logic.UserLogic.SignIn(email, password)
	if nil != err {
		return tools.Failed(400, err.Error())
	}
	token, err := logic.UserLogic.CheckLoginStatus(user.ID)
	if nil != err {
		repositories.MemoryStorageRepository.Del("user", token)
	}
	token = repositories.MemoryStorageRepository.Set("user", user)
	return tools.Success(map[string]interface{}{"token": token, "user": user})
}
