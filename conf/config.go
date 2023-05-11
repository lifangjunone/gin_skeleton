package conf

import "fmt"

type Config struct {
	App   *app   `toml:"app"`
	Mysql *mysql `toml:"mysql"`
	Mongo *mongo `toml:"mongo"`
}

type app struct {
	Name string `toml:"name" env:"APP_NAME"`
	Host string `toml:"host" env:"APP_HOST"`
	Port string `toml:"port" env:"APP_PORT"`
}

type mysql struct {
	Host     string `toml:"host" env:"MYSQL_HOST"`
	Port     string `toml:"port" env:"MYSQL_PORT"`
	UserName string `toml:"username" env:"MYSQL_USERNAME"`
	Password string `toml:"password" env:"MYSQL_PASSWORD"`
	Database string `toml:"database" env:"MYSQL_DATABASE"`
}

type mongo struct {
	Host     string `toml:"host" env:"MONGO_HOST"`
	Port     string `toml:"port" env:"MONGO_PORT"`
	UserName string `toml:"username" env:"MYSQL_USERNAME"`
	Password string `toml:"password" env:"MYSQL_PASSWORD"`
	Database string `toml:"database" env:"MYSQL_DATABASE"`
}

func (a *app) GetAddr() string {
	return fmt.Sprintf("%s:%s", a.Host, a.Port)
}

func newDefaultMysql() *mysql {
	return &mysql{
		Host:     "127.0.0.1",
		Port:     "3307",
		UserName: "root123",
		Password: "root123456",
		Database: "gin-skeleton",
	}
}

func newDefaultApp() *app {
	return &app{
		Host: "127.0.0.1",
		Port: "9090",
	}
}

func NewConfig() *Config {
	return &Config{
		App:   newDefaultApp(),
		Mysql: newDefaultMysql(),
	}
}
