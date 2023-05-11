package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
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

func InitConfig(configFilePath string) *Config {
	confObj := NewConfig()
	err := loadConfigFromFile(configFilePath, confObj)
	if err != nil {
		panic(err.Error())
	}
	return confObj
}
