package upload

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/MyBlog/global"
	"github.com/MyBlog/pkg/util"
)

type FileType int

const TypeImage FileType = iota + 1 // 1

// 返回加密后的文件名称
func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext) // 返回不带ext的文件名部分
	fileName = util.EncodeMD5(fileName)       // 加密

	return fileName + ext
}

// 获取文件后缀
func GetFileExt(name string) string {
	return path.Ext(name)
}

func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

// 检查保存目录是否存在
func CheckSavePath(dst string) bool {
	// 通过调用os.Stat方法获取文件的描述信息
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}
	return false
}

func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

// 检查文件权限是否够
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}
	return nil
}

func SaveFile(file *multipart.FileHeader, dst string) error {
	// 打开源地址的文件
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// 创建目标地址文件
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	// 文件内容拷贝
	_, err = io.Copy(out, src)
	return err
}
