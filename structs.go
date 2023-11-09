package config

import "github.com/spf13/viper"

// ConfigManager 配置管理器类型
type ConfigManager struct {
	// 配置相对路径
	configPath string
	//包内全局map，用于存储viper对象。key为模块名，value为viper对象。
	viperList map[string]*viper.Viper
}
