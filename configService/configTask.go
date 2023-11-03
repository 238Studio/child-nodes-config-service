package configService

import "github.com/spf13/viper"

// Start 启动模块服务
// 传入：启动参数
// 传出：无
func (conf *ConfigManager) Start() {
	viper.AddConfigPath(conf.configPath)
	viper.SetConfigType("json")
}

// Stop 中止模块服务
// 传入：无
// 传出：无
func (conf *ConfigManager) Stop() {

}

// GetApp 获取App
// 传入：无
// 传出：该模块App的指针
func (conf *ConfigManager) GetApp() *interface{} {
	var value interface{}
	value = conf
	return &value
}

// IsAlive 是否在服务
// 传入：无
// 传出：是否在服务
func (conf *ConfigManager) IsAlive() bool {
	return true
}
