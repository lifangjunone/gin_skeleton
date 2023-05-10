package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

func loadConfigFromFile(filePath string, confObj *Config) error {
	_, err := toml.DecodeFile(filePath, confObj)
	if err != nil {
		return err
	}
	return nil
}

func InitConfig(configFilePath string) {
	confObj := NewConfig()
	fmt.Println("2222222222")
	fmt.Println(configFilePath)
	fmt.Println("2222222222")
	err := loadConfigFromFile(configFilePath, confObj)
	if err != nil {
		panic(err.Error())
	}
}
