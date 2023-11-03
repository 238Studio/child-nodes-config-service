package configService_test

import (
	"github.com/UniversalRobotDriveTeam/child-nodes-basic/robotBasicAPI/configService"
	"testing"
)

// 测试利用管道调用配置模块
func TestInitConfigManager(t *testing.T) {

}
func TestSetConfig(t *testing.T) {
	// 初始化管道
	// 外部模块调用通道

	configManger := configService.InitConfigManager()
	//todo 测试管道
	configManger.Start()
	err := configManger.CreateConfigTable("test")
	if err != nil {
		t.Error(err)
	}

	var config = make(map[string]string)
	config["test"] = "test"
	config["test2"] = "test2"

	// 生成方法-参数对
	// 输入方法-参数对
	// 传入指针货的返回值
	// 管道阻塞
}
