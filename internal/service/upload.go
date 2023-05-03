package service

import (
	"errors"
	"mime/multipart"
	"os"

	"github.com/MyBlog/global"
	"github.com/MyBlog/pkg/upload"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	uploadSavePath := upload.GetSavePath()
	dst := uploadSavePath + "/" + fileName
	// 检查文件后缀是否在约定项中
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported.")
	}
	if upload.CheckSavePath(uploadSavePath) {
		err := upload.CreateSavePath(uploadSavePath, os.ModePerm)
		if err != nil {
			return nil, errors.New("failed to create save director.")
		}
	}
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceed maximum file limit.")
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions.")
	}
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}
	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}
