package routes

import (
	"github.com/kataras/iris/v12"
	"video_storage/logic"
	"video_storage/model"
	"video_storage/tools"
)

type DictionaryRoute struct {
	Ctx iris.Context
}

func (that *DictionaryRoute) GetVideoType() *model.ResponseBody {
	dictionaryType := tools.FormValue(that.Ctx, "dictionaryType")
	return tools.Success(logic.DictionaryLogic.GetDictionaryGroup(dictionaryType))
}

func (*DictionaryRoute) GetAllDictionary() *model.ResponseBody {
	return tools.Success(nil)
}