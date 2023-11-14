package config_test

import (
	"testing"

	config "github.com/UniversalRobotDriveTeam/child-nodes-config-service"
)

func TestDeleteConfig(t *testing.T) {
	configManger := config.InitConfigManager("../../configs")

	err := configManger.DeleteConfig("test", "test")
	if err != nil {
		t.Error(err)
	}
}

func TestConfigDelete(t *testing.T) {
	configManager := config.InitConfigManager("../../configs")
	err := configManager.DeleteConfigTable("test")
	if err != nil {
		t.Error(err)
	}
}
