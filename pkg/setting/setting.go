package setting

import "github.com/spf13/viper"

//针对配置文件的读取行为
type Setting struct {
	vp *viper.Viper
}

//初始化项目的配置的基础知识
func NewSetting() (*Setting, error) {
	vp := viper.New()
	//设置配置文件名称
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		 return nil , err
	}

	return &Setting{vp}, nil
}