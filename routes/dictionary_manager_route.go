package routes

import (
	"video_storage/logic"
	"video_storage/model"
	"video_storage/tools"

	"github.com/kataras/iris/v12"
)

type DictionaryManagerRoute struct {
	Ctx iris.Context
}

// 标签组添加标签
func (that *DictionaryManagerRoute) Put() *model.ResponseBody {
	name := tools.FormValue(that.Ctx, "name")
	data := tools.FormValue(that.Ctx, "data")
	groupID, err := tools.FormValueInt64(that.Ctx, "groupID")
	if nil != err {
		return tools.Failed(400, "参数错误")
	}
	if name == "" {
		return tools.Failed(400, "关键参数均不能为空")
	}
	if err := logic.DictionaryLogic.CreateDictionary(name, data, groupID); nil != err {
		return tools.Failed(400, err.Error())
	} else {
		return tools.Success(nil)
	}
}

// 删除子标签
func (that *DictionaryManagerRoute) Delete() *model.ResponseBody {
	dictionaryID, err := tools.FormValueInt(that.Ctx, "id")
	if nil != err {
		return tools.Failed(400, "参数类型错误")
	}
	err = logic.DictionaryLogic.DeleteDictionary(dictionaryID)
	if err != nil {
		return tools.Failed(500, err.Error())
	} else {
		return tools.Success(nil)
	}
}

// 创建标签组
func (that *DictionaryManagerRoute) PutGroup() *model.ResponseBody {
	name := tools.FormValue(that.Ctx, "name")
	groupType := tools.FormValue(that.Ctx, "groupType")
	if name == "" || groupType == "" {
		return tools.Failed(400, "关键参数均不能为空")
	}
	if err := logic.DictionaryLogic.CreateGroup(name, groupType); nil != err {
		return tools.Failed(500, err.Error())
	} else {
		return tools.Success(nil)
	}
}

func (that *DictionaryManagerRoute) PostGroup() *model.ResponseBody {
	groupID, err := tools.FormValueInt64(that.Ctx, "id")
	if nil != err {
		return tools.Failed(400, "对象不存在")
	}
	name := tools.FormValue(that.Ctx, "name")
	groupType := tools.FormValue(that.Ctx, "groupType")
	if name == "" || groupType == "" {
		return tools.Failed(400, "关键参数均不能为空")
	}
	if err := logic.DictionaryLogic.UpdateGroup(groupID, name, groupType); nil != err {
		return tools.Failed(500, err.Error())
	} else {
		return tools.Success(nil)
	}
}

// 删除标签组
func (that *DictionaryManagerRoute) DeleteGroup() *model.ResponseBody {
	groupID, err := tools.FormValueInt64(that.Ctx, "id")
	if nil != err {
		return tools.Failed(400, "对象不存在")
	}
	err = logic.DictionaryLogic.DeleteGroup(groupID)
	if err != nil {
		return tools.Failed(500, err.Error())
	} else {
		return tools.Success(nil)
	}
}

// 获取标签组集合
func (that *DictionaryManagerRoute) GetGroup() *model.ResponseBody {
	return tools.Success(logic.DictionaryLogic.GetDictionaryGroupList())
}
