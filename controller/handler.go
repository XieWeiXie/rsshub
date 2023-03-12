package controller

import "github.com/gin-gonic/gin"

const defaultXML = "application/rss+xml"

type Handler interface {
	ToRSSHandler(ctx *gin.Context)
	Describe() string
}

var cacheHandler map[string]Handler

func init() {
	cacheHandler = make(map[string]Handler)
}

func NewHandler(key string) Handler {
	v, ok := cacheHandler[key]
	if ok {
		return v
	}
	return nil
}
