package routes

import (
	"video_storage/logic"
	"video_storage/model"
	"video_storage/tools"

	"github.com/kataras/iris/v12"
)

type ObjectRoute struct {
	Ctx iris.Context
}

func (that *ObjectRoute) Put() *model.ResponseBody {
	metadataID, err := tools.FormValueInt64(that.Ctx, "id")
	if err != nil {
		return tools.Failed(400, "源数据id无效")
	}
	fileMetadata, err := logic.FileMetadataLogic.FindByID(metadataID)
	if fileMetadata.IsDel || nil != err {
		return tools.Failed(404, "源数据已被删除")
	}
	_, err = logic.ObjectLogic.FindByID(metadataID)
	if nil == err {
		return tools.Failed(400, "对象已入库")
	}
	err = logic.ObjectLogic.SaveObject(fileMetadata)
	if nil != err {
		return tools.Failed(500, "对象入库失败")
	}
	isVideo, _ := tools.FormValueBool(that.Ctx, "isVideo")
	if isVideo {
		width, _ := tools.FormValueInt64(that.Ctx, "width")
		height, _ := tools.FormValueInt64(that.Ctx, "height")
		codecName := tools.FormValue(that.Ctx, "codec_name")
		displayAspectRatio := tools.FormValue(that.Ctx, "display_aspect_ratio")
		codecLongName := tools.FormValue(that.Ctx, "codec_long_name")
		duration := tools.FormValue(that.Ctx, "duration")
		videoDescribe := &model.VideoDescribe{}
		videoDescribe.Width = width
		videoDescribe.MetadataID = fileMetadata.ID
		videoDescribe.Height = height
		videoDescribe.DisplayAspectRatio = displayAspectRatio
		videoDescribe.CodecName = codecName
		videoDescribe.Duration = duration
		videoDescribe.CodecLongName = codecLongName
		_ = logic.VideoLogic.SaveDescribe(videoDescribe)
	}

	return tools.Success(nil)
}
