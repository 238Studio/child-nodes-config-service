package configService_test

import (
	"github.com/UniversalRobotDriveTeam/child-nodes-basic/robotBasicAPI/configService"
	"testing"
)

func TestDeleteConfig(t *testing.T) {
	configManger := configService.InitConfigManager()

	err := configManger.DeleteConfig("test", "test")
	if err != nil {
		t.Error(err)
	}
}

func TestConfigDelete(t *testing.T) {
	configManager := configService.InitConfigManager()
	err := configManager.DeleteConfigTable("test")
	if err != nil {
		t.Error(err)
	}
}
