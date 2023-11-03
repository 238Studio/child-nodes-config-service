package configService

// ConfigManager 配置管理器类型
type ConfigManager struct {
	// 配置相对路径
	configPath string
	// 外部模块调用通道
	ConfigCallChannel *chan map[string]interface{}
}
