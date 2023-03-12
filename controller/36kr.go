package controller

import (
	"github.com/XieWeiXie/rsshub/pkg/hub"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Kr36 struct {
}

func (k Kr36) Describe() string {
	describe := "36kr"
	cacheHandler[describe] = k
	return describe
}

func (k Kr36) ToRSSHandler(ctx *gin.Context) {
	rule := hub.Rule{
		Host:            "https://36kr.com",
		HostTitle:       "36kr",
		HostDescription: "36氪_让一部分人先看到未来",
		TargetURl:       "https://36kr.com/",
		ListContainers:  "#app .main-right .kr-home-main .kr-home-flow .kr-home-flow-list .kr-flow-article-item",
		Title:           ".article-item-title",
		URL:             ".article-item-title",
		Date:            ".kr-flow-bar-time",
		Contents:        ".article-item-description",
		Author:          ".kr-flow-bar-author",
		Description:     ".article-item-description",
	}
	toRes := hub.ToResponseHTTP{}
	doc, _ := toRes.Response(rule.TargetURl)
	feeds := hub.NewFeeds()
	f, _ := feeds.ToFeed(doc, rule)
	rss, _ := f.ToRss()
	ctx.Data(http.StatusOK, defaultXML, []byte(rss))
}
