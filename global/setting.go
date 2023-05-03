package global

import (
	"github.com/MyBlog/pkg/logger"
	"github.com/MyBlog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS

	Logger *logger.Logger

	JWTSetting *setting.JWTSettingS

	EmailSetting *setting.EmailSettingS
)
