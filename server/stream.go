package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go/jetstream"
	"log"
	"net/http"
)

type Stream struct{}

func (stm *Stream) StreamIndex(ctx *gin.Context) {
	streams := make([]map[string]any, 0)

	streamList := jetStreams[selectUrl].ListStreams(ctx)
	for s := range streamList.Info() {
		streams = append(streams, map[string]any{
			"name":           s.Config.Name,
			"subjects":       s.Config.Subjects,
			"num_replicas":   s.Config.Replicas,
			"messages":       s.State.Msgs,
			"bytes":          s.State.Bytes,
			"consumer_count": s.State.Consumers,
			"create_at":      s.Created,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":     "ok",
		"status_code": http.StatusOK,
		"data":        streams,
	})
}

func (stm *Stream) StreamStore(ctx *gin.Context) {
	name := ctx.PostForm("name")
	subjects := ctx.PostFormArray("subjects[]")

	if len(subjects) == 0 || name == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message":     "主题名称或者主题必填",
			"status_code": http.StatusUnprocessableEntity,
		})
		return
	}

	_, err := jetStreams[selectUrl].CreateStream(ctx, jetstream.StreamConfig{
		Name:     name,
		Subjects: subjects,
	})
	if err != nil {
		log.Printf("创建 stream 失败: %v", err)
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
			"name":     name,
			"subjects": subjects,
		},
	})
}

func (stm *Stream) StreamShow(ctx *gin.Context) {
	name := ctx.Param("stream")

	stream, err := jetStreams[selectUrl].Stream(ctx, name)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message":     "not found",
			"status_code": http.StatusNotFound,
		})
		return
	}

	info, err := stream.Info(ctx)
	if err != nil {
		log.Printf("获取 stream 信息失败: %v", err)
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
			"config":  info.Config,
			"state":   info.State,
			"created": info.Created,
		},
	})
}

func (stm *Stream) StreamUpdate(ctx *gin.Context) {
	name := ctx.Param("stream")
	subjects := ctx.PostFormArray("subjects[]")

	if len(subjects) == 0 || name == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message":     "主题名称或者主题必填",
			"status_code": http.StatusUnprocessableEntity,
		})
		return
	}

	_, err := jetStreams[selectUrl].UpdateStream(ctx, jetstream.StreamConfig{
		Name:     name,
		Subjects: subjects,
	})
	if err != nil {
		log.Printf("更新 stream 失败: %v", err)
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
			"name":     name,
			"subjects": subjects,
		},
	})
}

func (stm *Stream) StreamDelete(ctx *gin.Context) {
	name := ctx.Param("stream")

	err := jetStreams[selectUrl].DeleteStream(ctx, name)
	if err != nil {
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
			"name": name,
		},
	})
}
