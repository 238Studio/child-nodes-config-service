package config

import "github.com/spf13/viper"

// 配置相对路径
//const configPath = "./configs/test.json"
//const configPath = "../../configs" //测试路径

// InitConfigManager 初始化配置管理器
// 传入:配置文件路径
// 传出:配置文件管理结构
func InitConfigManager(configPath string) *ConfigManager {
	c := ConfigManager{
		configPath: configPath,
	}

	c.viperList = make(map[string]*viper.Viper)
	return &c
}

// InitModuleConfig 初始化模块配置
// 传入:模块名
// 传出:无
func (conf *ConfigManager) InitModuleConfig(moduleName string) {
	//初始化模块viper对象
	conf.viperList[moduleName] = viper.New()
	conf.viperList[moduleName].SetConfigType("json")
	conf.viperList[moduleName].SetConfigName(moduleName)
	conf.viperList[moduleName].AddConfigPath(conf.configPath + "/")
}
