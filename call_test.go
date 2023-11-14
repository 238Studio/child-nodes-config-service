package config_test

import (
	"testing"

	config "github.com/UniversalRobotDriveTeam/child-nodes-config-service"
)

// 测试利用管道调用配置模块
func TestInitConfigManager(t *testing.T) {

}
func TestSetConfig(t *testing.T) {
	// 初始化管道
	// 外部模块调用通道

	configManger := config.InitConfigManager("./configs")
	//todo 测试管道
	configManger.InitModuleConfig("test")
	err := configManger.CreateConfigTable("test")
	if err != nil {
		t.Error(err)
	}

	var config = make(map[string]string)
	config["test"] = "test"
	config["test2"] = "test2"

	configManger.SetConfig("test", config)

	configManger.DeleteConfig("test", "test2")

	var newConfig = make(map[string]string)
	newConfig["test3"] = "test"
	newConfig["test"] = "test2"
	configManger.SetConfig("test", newConfig)

	item, _ := configManger.ReadConfig("test", "test")
	t.Log(item)

	//configManger.DeleteConfigTable("test")
}
