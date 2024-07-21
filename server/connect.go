package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"net/http"
)

func TestConnect(c *gin.Context) {
	host := c.Query("host")
	port := c.Query("port")

	if host == "" || port == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message":     "host 或者 port 未传递",
			"status_code": http.StatusUnprocessableEntity,
		})
		return
	}

	nc, err := nats.Connect(fmt.Sprintf("nats://%s:%s", host, port))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":     "地址错误或 nats 服务未启动",
			"status_code": http.StatusBadRequest,
		})
		return
	}
	defer nc.Close()

	c.JSON(http.StatusOK, gin.H{
		"message":     "连接成功",
		"status_code": http.StatusOK,
	})
}
