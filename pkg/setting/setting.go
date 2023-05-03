package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
) 

//配置管理
type Setting struct {
	vp *viper.Viper
}

// 初始化本项目配置的基础属性
func NewSetting(configs ...string) (*Setting, error) {
	vp := viper.New()
	// 设定配置文件的名称
	vp.SetConfigName("config")
	for _, config := range configs {
		if config != "" {
			vp.AddConfigPath(config)
		}
	}
	// 配置路径
	vp.AddConfigPath("configs/")
	// 配置文件的类型
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	s := &Setting{vp}
	s.WatchSettingChange()
	return s, nil
}

func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = s.ReloadAllSection()
		})
	}()
}
