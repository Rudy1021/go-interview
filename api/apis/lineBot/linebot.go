package lineBot

import (
	"fmt"
	"go-interview/api/apis/collections"
	"go-interview/api/config"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var bot *linebot.Client

func Init() *linebot.Client {
	bot, err := linebot.New(config.Getstr("lineBot.channel_secret"), config.Getstr("lineBot.channel_access_token"))
	if err != nil {
		fmt.Println(err)
	}
	return bot
}
func Mess(c *gin.Context) {
	bot := Init()
	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			log.Print(err)
		}
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if message.Text == "show" {
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(collections.UserText_r_one(event.Source.UserID))).Do()

				} else if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					log.Print(err)
				} else {
					collections.UserText_c(event.Source.UserID, message.Text)
				}
			}
		}
	}
}
