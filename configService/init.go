package configService

import "github.com/spf13/viper"

// 配置相对路径
//const configPath = "./configs/test.json"

const configPath = "../../configs" //测试路径

var viperList = make(map[string]*viper.Viper) //包内全局map，用于存储viper对象。key为模块名，value为viper对象。

func InitConfigManager() *ConfigManager {
	c := ConfigManager{configPath: configPath}
	viper.AddConfigPath(configPath)
	viper.SetConfigType("json")
	return &c
}
