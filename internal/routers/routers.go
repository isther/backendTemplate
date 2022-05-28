package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()
	router.Static("/static", "static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", index)
	router.GET("/ping", ping)

	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin":  "password",
		"admin1": "password",
	}))
	{
		authorized.GET("/secrets", authSecret)
	}
	return router
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "Hello",
	})
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func authSecret(c *gin.Context) {
	var secrets = gin.H{
		"admin": gin.H{"email": "admin@admin.com", "phone": "123433"},
	}

	user := c.MustGet(gin.AuthUserKey).(string)
	if secret, ok := secrets[user]; ok {
		c.JSON(http.StatusOK, gin.H{
			"user":   user,
			"secret": secret,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"user":   user,
			"secret": "NO SECRET :<",
		})
	}
}
