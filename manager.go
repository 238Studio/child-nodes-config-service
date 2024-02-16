package config

import (
	"errors"
	"github.com/238Studio/child-nodes-error-manager/errpack"
	"os"
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
			return errpack.NewError(errpack.CommonException, errpack.Config, err)
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

	return errpack.NewError(errpack.CommonException, errpack.Config, err)
}

// ReadConfig 某个模块根据模块-配置名 读取相应配置
// 传入：该模块名称，该模块项目名
// 传出：该配置项目
func (conf *ConfigManager) ReadConfig(module string, configItem string) (string, error) {
	v, isExist := conf.viperList[module]
	if !isExist {
		return "", errpack.NewError(errpack.CommonException, errpack.Config, errors.New(module+"模块未在配置文件管理器内注册"))
	}

	err := v.ReadInConfig()
	item, err := v.GetString(configItem), errpack.NewError(errpack.CommonException, errpack.Config, err)
	return item, nil
}

// DeleteConfig 删除该项配置
// 传入：模块名，该模块项目名
// 传出：无
func (conf *ConfigManager) DeleteConfig(module string, configItem string) error {
	v, isExist := conf.viperList[module]
	if !isExist {
		return errpack.NewError(errpack.CommonException, errpack.Config, errors.New(module+"模块未在配置文件管理器内注册"))
	}

	v.Set(configItem, nil)
	err := v.WriteConfig()
	return errpack.NewError(errpack.CommonException, errpack.Config, err)
}

// SetConfig 设置该项配置
// 传入：模块名，该模块项目名，参数
// 传出：无
func (conf *ConfigManager) SetConfig(module string, configItems map[string]string) error {
	v, isExist := conf.viperList[module]
	if !isExist {
		return errpack.NewError(errpack.CommonException, errpack.Config, errors.New(module+"模块未在配置文件管理器内注册"))
	}

	for configItem, item := range configItems {
		v.Set(configItem, item)
	}

	err := v.WriteConfig()
	return errpack.NewError(errpack.CommonException, errpack.Config, err)
}
