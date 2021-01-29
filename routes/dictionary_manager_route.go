package routes

import (
	"github.com/kataras/iris/v12"
	"video_storage/logic"
	"video_storage/model"
	"video_storage/tools"
)

type DictionaryManagerRoute struct {
	Ctx iris.Context
}

// 创建标签组
func (that *DictionaryRoute) PutDictionaryGroup() *model.ResponseBody {
	name := tools.FormValue(that.Ctx, "name")
	groupType := tools.FormValue(that.Ctx, "groupType")
	if "" == name || "" == groupType {
		return tools.Failed(400, "关键参数均不能为空")
	}
	if err := logic.DictionaryLogic.CreateGroup(name, groupType); nil != err {
		return tools.Failed(400, err.Error())
	} else {
		return tools.Success(nil)
	}
}