package main

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	Node
	Auth
	Stream
	Consumer
}

var selectUrl string

func main() {
	app := new(App)

	r := gin.Default()
	r.Use(cors())
	r.GET("/connect", TestConnect)

	r.GET("/nodes", app.NodeIndex)
	r.POST("/nodes", app.NodeStore)

	r.POST("/auths", app.AuthStore)

	auth := r.Group("")
	auth.Use(authorize())
	{
		auth.GET("/streams", app.StreamIndex)
		auth.POST("/streams", app.StreamStore)
		auth.GET("/streams/:stream", app.StreamShow)
		auth.PUT("/streams/:stream", app.StreamUpdate)
		auth.DELETE("/streams/:stream", app.StreamDelete)

		auth.GET("/consumers", app.ConsumerIndex)
		auth.POST("/consumers", app.ConsumerStore)
		auth.GET("/consumers/:consumer", app.ConsumerShow)
		auth.PUT("/consumers/:consumer", app.ConsumerUpdate)
		auth.DELETE("/consumers/:consumer", app.ConsumerDelete)
	}

	_ = r.Run(":8080")
}
