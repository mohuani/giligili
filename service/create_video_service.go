package service

import (
	"giligili/model"
	"giligili/serializer"
)

// CreateVideoService 视频创建服务
type CreateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info  string `form:"info" json:"info" binding:"required,max=200"`
}

// Create 视频创建
func (service *CreateVideoService) Create() serializer.Response {
	video := model.Video{
		Title: service.Title,
		Info:  service.Info,
	}

	err := model.DB.Create(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  50001,
			Data:  nil,
			Msg:   "视频保存失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
