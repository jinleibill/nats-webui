package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go/jetstream"
	"log"
	"net/http"
)

type Consumer struct{}

func (con *Consumer) ConsumerIndex(ctx *gin.Context) {
	consumers := make([]map[string]any, 0)

	names := jetStreams[selectUrl].StreamNames(ctx)
	for name := range names.Name() {
		stream, err := jetStreams[selectUrl].Stream(ctx, name)
		if err != nil {
			continue
		}
		for c := range stream.ListConsumers(ctx).Info() {
			consumers = append(consumers, map[string]any{
				"name":            c.Name,
				"durable_name":    c.Config.Durable,
				"stream_name":     c.Stream,
				"num_ack_pending": c.NumAckPending,
				"num_redelivered": c.NumRedelivered,
				"num_waiting":     c.NumWaiting,
				"num_pending":     c.NumPending,
				"create_at":       c.Created,
			})
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":     "ok",
		"status_code": http.StatusOK,
		"data":        consumers,
	})
}

func (con *Consumer) ConsumerStore(ctx *gin.Context) {
	streamName := ctx.PostForm("stream_name")
	name := ctx.PostForm("name")
	durableName := ctx.PostForm("durable_name")

	if streamName == "" || durableName == "" || name == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message":     "stream 名称, 消费者名称或者 durable 必填",
			"status_code": http.StatusUnprocessableEntity,
		})
		return
	}

	_, err := jetStreams[selectUrl].CreateConsumer(ctx, streamName, jetstream.ConsumerConfig{
		Name:      name,
		Durable:   durableName,
		AckPolicy: jetstream.AckExplicitPolicy,
	})
	if err != nil {
		log.Printf("创建 Consumer 失败: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":     "创建失败",
			"status_code": http.StatusBadRequest,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":     "ok",
		"status_code": http.StatusCreated,
		"data": map[string]any{
			"stream_name":  streamName,
			"name":         name,
			"durable_name": durableName,
		},
	})
}

func (con *Consumer) ConsumerShow(ctx *gin.Context) {
	name := ctx.Param("consumer")
	streamName := ctx.Query("stream_name")

	consumer, err := jetStreams[selectUrl].Consumer(ctx, streamName, name)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message":     "not found",
			"status_code": http.StatusNotFound,
		})
		return
	}

	info, err := consumer.Info(ctx)
	if err != nil {
		log.Printf("获取 Consumer 信息失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":     "内部错误",
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":     "ok",
		"status_code": http.StatusOK,
		"data": map[string]any{
			"info": info,
		},
	})
}

func (con *Consumer) ConsumerUpdate(ctx *gin.Context) {
	streamName := ctx.PostForm("stream_name")
	name := ctx.Param("consumer")
	durableName := ctx.PostForm("durable_name")

	if streamName == "" || durableName == "" || name == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message":     "stream 名称, 消费者名称或者 durable 必填",
			"status_code": http.StatusUnprocessableEntity,
		})
		return
	}

	_, err := jetStreams[selectUrl].UpdateConsumer(ctx, streamName, jetstream.ConsumerConfig{
		Name:      name,
		Durable:   durableName,
		AckPolicy: jetstream.AckExplicitPolicy,
	})
	if err != nil {
		log.Printf("更新 Consumer 失败: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":     "更新失败",
			"status_code": http.StatusBadRequest,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":     "ok",
		"status_code": http.StatusOK,
		"data": map[string]any{
			"stream_name":  streamName,
			"name":         name,
			"durable_name": durableName,
		},
	})
}

func (con *Consumer) ConsumerDelete(ctx *gin.Context) {
	name := ctx.Param("consumer")
	streamName := ctx.Query("stream_name")

	err := jetStreams[selectUrl].DeleteConsumer(ctx, streamName, name)
	if err != nil {
		log.Printf("删除consumer 失败: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":     "删除失败",
			"status_code": http.StatusBadRequest,
		})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{
		"message":     "ok",
		"status_code": http.StatusNoContent,
		"data": map[string]any{
			"name":        name,
			"stream_name": streamName,
		},
	})
}
