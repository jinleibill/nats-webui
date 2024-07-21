package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"log"
	"net/http"
)

var connects = make(map[string]*nats.Conn)
var jetStreams = make(map[string]jetstream.JetStream)

type Auth struct{}

func (a *Auth) AuthStore(ctx *gin.Context) {
	url := ctx.PostForm("url")

	if url == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message":     "url 未传递",
			"status_code": http.StatusUnprocessableEntity,
		})
		return
	}

	nc, err := nats.Connect(fmt.Sprintf("nats://%s", url))
	if err != nil {
		log.Printf("连接 nats 服务失败, 错误: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":     "地址错误或 nats 服务未启动",
			"status_code": http.StatusBadRequest,
		})
		return
	}
	connects[url] = nc

	js, err := jetstream.New(nc)
	if err != nil {
		log.Printf("创建 jetstream 实例失败, 错误: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":     "内部错误",
			"status_code": http.StatusInternalServerError,
		})
		return
	}
	jetStreams[url] = js

	selectUrl = url

	ctx.JSON(http.StatusCreated, gin.H{
		"message":     "ok",
		"status_code": http.StatusCreated,
		"data":        map[string]string{"url": url},
	})
}
