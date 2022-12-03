package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	fmt.Println(Getstr("lineBot.channel_secret"))
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
