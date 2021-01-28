package routes

import (
	"github.com/kataras/iris/v12"
	"video_storage/model"
	"video_storage/tools"
)

type DictionaryRoute struct {
	Ctx iris.Context
}

func (*DictionaryRoute) GetVideoType() *model.ResponseBody {
	return tools.Success([]string{"热度推荐", "电影", "电视剧", "动漫"})
}