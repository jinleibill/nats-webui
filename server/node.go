package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var db = make([]*Node, 0)

type Node struct {
	Name     string    `json:"name"`
	Host     string    `json:"host"`
	Port     string    `json:"port"`
	CreateAt time.Time `json:"create_at"`
}

func (n *Node) NodeIndex(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":     "ok",
		"status_code": http.StatusOK,
		"data":        db,
	})
}

func (n *Node) NodeStore(ctx *gin.Context) {
	name := ctx.PostForm("name")
	host := ctx.PostForm("host")
	port := ctx.PostForm("port")

	if host == "" || port == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message":     "host 或者 port 未传递",
			"status_code": http.StatusUnprocessableEntity,
		})
		return
	}

	newNode := &Node{
		Name:     name,
		Host:     host,
		Port:     port,
		CreateAt: time.Now(),
	}

	db = append(db, newNode)

	ctx.JSON(http.StatusCreated, gin.H{
		"message":     "ok",
		"status_code": http.StatusCreated,
		"data":        newNode,
	})
}
