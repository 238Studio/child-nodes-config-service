package config_test

import (
	"testing"

	config "github.com/238Studio/child-nodes-config-service"
)

func TestConfig(t *testing.T) {
	// 初始化管道
	// 外部模块调用通道

	configManger := config.InitConfigManager("./configs")
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
	if err != nil { //注释18~21行
		t.Log("panic recover")
		t.Error(err)
	}

	configManger.DeleteConfig("test", "test2")

	var newConfig = make(map[string]string)
	newConfig["test3"] = "test"
	newConfig["test"] = "test2"
	configManger.SetConfig("test", newConfig)

	item, _ := configManger.ReadConfig("test", "test")
	t.Log(item)

	//configManger.DeleteConfigTable("test")
}
