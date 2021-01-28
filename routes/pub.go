package routes

import (
	"github.com/kataras/iris/v12"
	"video_storage/config"
	"video_storage/model"
)

type PubRoute struct {
	Ctx iris.Context
}

// 获取服务器目前版本
func (*PubRoute) GetVersion() *model.ResponseBody {
	return &model.ResponseBody{
		Code:    200,
		Message: config.Instance.Common.Version,
	}
}
