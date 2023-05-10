package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
)

var (
	ConfObj *Config
)

func loadConfigFromFile(filePath string, confObj *Config) error {
	cur, _ := os.Getwd()
	filename := filepath.Join(cur, "conf/dev/dev.toml")
	fmt.Println(filename)
	_, err := toml.DecodeFile(filePath, confObj)
	if err != nil {
		return err
	}
	return nil
}

func InitConfig(configFilePath string) {
	confObj := NewConfig()
	err := loadConfigFromFile(configFilePath, confObj)
	if err != nil {
		panic(err.Error())
	}
	ConfObj = confObj
}
