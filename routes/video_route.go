package routes

import (
	"video_storage/logic"
	"video_storage/model"
	"video_storage/tools"

	"github.com/kataras/iris/v12"
)

type VideoRoute struct {
	Ctx iris.Context
}

func (that *VideoRoute) GetAnalysisFile() *model.ResponseBody {
	id, err := tools.FormValueInt64(that.Ctx, "id")
	if nil != err {
		return tools.Failed(400, "参数错误")
	}
	file, err := logic.FileMetadataLogic.FindByID(id)
	if nil != err {
		return tools.Failed(404, err.Error())
	}
	if file.IsDel {
		return tools.Failed(404, "对象已经被删除")
	}
	info, err := logic.VideoLogic.Ffprobe(file.FullPath)
	if nil != err {
		return tools.Failed(403, err.Error())
	}

	return tools.Success(info)
}

func (that *VideoRoute) GetPlayer() *model.ResponseBody {
	id, err := tools.FormValueInt64(that.Ctx, "id")
	if nil != err {
		return tools.Failed(400, err.Error())
	}
	if stream, err := logic.VideoLogic.PlayVideo(id); nil != err {
		return tools.Failed(400, err.Error())
	} else {
		return tools.Success(stream)
	}
}

// 上新
// func (that *VideoRoute) PostFfmpeg() *model.ResponseBody {
// 	videoPath := tools.FormValue(that.Ctx, "path")
// 	video, err := logic.VideoLogic.FFmpeg(videoPath)
// 	if nil != err {
// 		return tools.Failed(400, err.Error())
// 	} else {
// 		return tools.Success(video)
// 	}
// }

// 视频目录
func (that *VideoRoute) GetList() *model.ResponseBody {
	videoType := tools.FormValueDefault(that.Ctx, "type", "hot")
	page := tools.FormValueIntDefault(that.Ctx, "page", 0)
	count := tools.FormValueIntDefault(that.Ctx, "count", 10)
	page -= 1
	if page < 0 {
		page = 0
	}
	return tools.Success(logic.VideoLogic.FindVideoByTarget(videoType, page, count))
}
