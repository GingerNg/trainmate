package initialize

import (
	conf "trainmate/conf"
)

var (
	config *conf.Config
)

func InitConf() {
	config = conf.GetGlobalConfig()
	// fmt.Println(config)
}
