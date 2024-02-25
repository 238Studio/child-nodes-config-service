package config_test

import (
	"testing"

	config "github.com/238Studio/child-nodes-config-service"
)

func TestSetConfig(t *testing.T) {
	// 初始化管道
	// 外部模块调用通道

	configManger := config.InitConfigManager("./configs")
	//todo 测试管道
	err := configManger.InitModuleConfig("test")
	if err != nil {
		t.Error(err)
	}

	err = configManger.CreateConfigTable("test")
	if err != nil {
		t.Error(err)
	}

	var config = make(map[string]string)
	config["test"] = "test"
	config["test2"] = "test2"

	err = configManger.SetConfig("test", config)
	if err != nil {
		t.Error(err)
	}

	err = configManger.DeleteConfig("test", "test2")
	if err != nil {
		t.Error(err)
	}

	var newConfig = make(map[string]string)
	newConfig["test3"] = "test"
	newConfig["test"] = "test2"
	err = configManger.SetConfig("test", newConfig)
	if err != nil {
		t.Error(err)
	}

	item, _ := configManger.ReadConfig("test", "test")
	t.Log(item)
	item, _ = configManger.ReadConfig("test", "test3")
	t.Log(item)

	err = configManger.DeleteConfigTable("test")
	if err != nil {
		t.Error(err)
	}
}
