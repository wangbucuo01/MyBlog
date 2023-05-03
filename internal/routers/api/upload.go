package api

import (
	"github.com/gin-gonic/gin"
	"github.com/MyBlog/global"
	"github.com/MyBlog/internal/service"
	"github.com/MyBlog/pkg/app"
	"github.com/MyBlog/pkg/convert"
	"github.com/MyBlog/pkg/errcode"
	"github.com/MyBlog/pkg/upload"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if err != nil {
		errRsp := errcode.InvalidParams.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf("svc.UploadFIle err:%v", err)
		errRsp := errcode.ErrorUploadFileFail.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}
	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
