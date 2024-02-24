package config

import (
	"errors"
	"os"

	"github.com/238Studio/child-nodes-error-manager/errpack"
)

// CreateConfigTable 创建一个配置表单
// 传入：模块名
// 传出：无
func (conf *ConfigManager) CreateConfigTable(module string) error {
	f := conf.configPath + "/" + module + ".json"
	if _, err := os.Stat(f); err != nil {
		if os.IsNotExist(err) {
			file, _ := os.Create(f)
			err := file.Close() //关闭文件
			if err != nil {
				return errpack.NewError(errpack.CommonException, errpack.Config, err)
			}
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
	if err != nil {
		return errpack.NewError(errpack.CommonException, errpack.Config, err)
	}

	//从viperList中删除该模块
	delete(conf.viperList, module)

	return nil
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
	if err != nil {
		return "", errpack.NewError(errpack.CommonException, errpack.Config, err)
	}

	item := v.GetString(configItem)
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
	if err != nil {
		return errpack.NewError(errpack.CommonException, errpack.Config, err)
	}

	return nil
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
	if err != nil {
		return errpack.NewError(errpack.CommonException, errpack.Config, err)
	}

	return nil
}
