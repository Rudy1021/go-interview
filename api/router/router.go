package router

import (
	. "go-interview/api/apis/collections"
	"go-interview/api/apis/lineBot"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(CORS())
	router.POST("/callback", lineBot.Mess)
	/*-----------------------userText-----------------------*/
	router.GET("/userText/Select/:id", UserText_r_one)
	router.POST("/userText/Upload", UserText_c)
	return router
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "accept, authorization, content-type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
