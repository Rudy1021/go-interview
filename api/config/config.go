package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigFile("../../config.conf.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig()failed, err:%v\n", err)
		return
	}
	viper.WatchConfig()
	return
}

func Getstr(name string) string {
	return viper.GetString(name)
}
