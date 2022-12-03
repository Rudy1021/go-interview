package linebot

import (
	"fmt"
	"go-interview/api/config"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func init() {
	bot, err := linebot.New(config.Getstr("lineBot.channel_secret"), config.Getstr("lineBot.channel_access_token"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("bot: %v\n", bot)
}
