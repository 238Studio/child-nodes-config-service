package config

import (
	"errors"
	"os"

	_const "github.com/238Studio/child-nodes-assist/const"
	"github.com/238Studio/child-nodes-assist/util"
)

// CreateConfigTable 创建一个配置表单
// 传入：模块名
// 传出：无
func (conf *ConfigManager) CreateConfigTable(module string) error {
	file := conf.configPath + "/" + module + ".json"
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			_, _ = os.Create(file)
		} else {
			return util.NewError(_const.CommonException, _const.CreateConfigTableError, err)
		}
	} else {
		return nil
	}
	return nil
}

// DeleteConfigTable 删除一个表单
// 传入：模块名
// 传出：无
func (conf *ConfigManager) DeleteConfigTable(module string) error {
	err := os.Remove(conf.configPath + "/" + module + ".json")

	//从viperList中删除该模块
	delete(conf.viperList, module)

	return util.NewError(_const.CommonException, _const.DeleteConfigTableError, err)
}

// ReadConfig 某个模块根据模块-配置名 读取相应配置
// 传入：该模块名称，该模块项目名
// 传出：该配置项目
func (conf *ConfigManager) ReadConfig(module string, configItem string) (string, error) {
	v, isExist := conf.viperList[module]
	if !isExist {
		return "", util.NewError(_const.CommonException, _const.ConfigModuleNotExist, errors.New(module+"模块不存在"))
	}

	err := v.ReadInConfig()
	item, err := v.GetString(configItem), util.NewError(_const.CommonException, _const.ReadConfigError, err)
	return item, nil
}

// DeleteConfig 删除该项配置
// 传入：模块名，该模块项目名
// 传出：无
func (conf *ConfigManager) DeleteConfig(module string, configItem string) error {
	v, isExist := conf.viperList[module]
	if !isExist {
		return util.NewError(_const.CommonException, _const.ConfigModuleNotExist, errors.New(module+"模块不存在"))
	}

	v.Set(configItem, nil)
	err := v.WriteConfig()
	return util.NewError(_const.CommonException, _const.WriteConfigError, err)
}

// SetConfig 设置该项配置
// 传入：模块名，该模块项目名，参数
// 传出：无
func (conf *ConfigManager) SetConfig(module string, configItems map[string]string) error {
	v, isExist := conf.viperList[module]
	if !isExist {
		return util.NewError(_const.CommonException, _const.ConfigModuleNotExist, errors.New(module+"模块不存在"))
	}

	for configItem, item := range configItems {
		v.Set(configItem, item)
	}

	err := v.WriteConfig()
	return util.NewError(_const.CommonException, _const.WriteConfigError, err)
}
