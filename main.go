package main

import (
	"github.com/gin-gonic/gin"
	"github.com/isther/web-base/conf"
	log "github.com/sirupsen/logrus"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	log.Warn("test")
	r.Run(conf.Server.Listen)
}
