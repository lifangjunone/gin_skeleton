package global

import "gin_skeleton/conf"

var (
	Service        = "gin-skeleton"
	Version        = "v1.0.0"
	ConfigFilePath = "../conf/env/dev.toml"
	ConfigObj      *conf.Config
)
