package global

import "gin_skeleton/conf"

var (
	Service        = "gin-skeleton"
	Version        = "v1.0.0"
	ConfigFilePath string
	ConfigObj      *conf.Config
)
