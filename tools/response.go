package tools

import "video_storage/model"

func Failed(code int, message string) *model.ResponseBody {
	return &model.ResponseBody{
		Code:    code,
		Message: message,
		Context: nil,
	}
}

func Success(context interface{}) *model.ResponseBody {
	return &model.ResponseBody{
		Code:    200,
		Message: "",
		Context: context,
	}
}
