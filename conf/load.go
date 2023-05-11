package conf

import (
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
)

func loadConfigFromFile(filePath string, confObj *Config) error {
	var filename string
	if filePath == "" {
		cur, _ := os.Getwd()
		filename = filepath.Join(cur, "conf/dev/dev.toml")
	} else {
		filename = filePath
	}
	_, err := toml.DecodeFile(filename, confObj)
	if err != nil {
		return err
	}
	return nil
}

func InitConfig(configFilePath string) *Config {
	confObj := NewConfig()
	err := loadConfigFromFile(configFilePath, confObj)
	if err != nil {
		panic(err.Error())
	}
	return confObj
}
