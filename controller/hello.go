package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Hello struct {
}

func (h Hello) Describe() string {
	describe := "hello"
	cacheHandler[describe] = h
	return describe
}

func (h Hello) ToRSSHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data":      "hi",
		"timestamp": time.Now().Unix(),
		"date":      time.Now().Format("2006-01-02 15:04:05"),
	})
}
