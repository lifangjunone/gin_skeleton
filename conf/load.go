package conf

import (
	"gin_skeleton/global"
	"github.com/BurntSushi/toml"
)

func loadConfigFromFile(filePath string, confObj *Config) error {
	_, err := toml.DecodeFile(filePath, confObj)
	if err != nil {
		return err
	}
	return nil
}

func InitConfig() {
	confObj := NewConfig()
	err := loadConfigFromFile(global.ConfigFilePath, confObj)
	if err != nil {
		panic(err.Error())
	}
	global.ConfigObj = confObj
}
