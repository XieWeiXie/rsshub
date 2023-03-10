package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/XieWeiXie/rsshub"
	"strings"
)

func main() {
	NewDaYu()
}

func New36kr() {
	rule := rsshub.Rule{
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

	toRes := rsshub.ToResponseRod{}
	doc, _ := toRes.Response(rule.TargetURl)

	feeds := rsshub.NewFeeds()
	f, _ := feeds.ToFeed(doc, rule)
	fmt.Println(f.ToRss())
}

func New36krNews() {
	rule := rsshub.Rule{
		Host:            "https://36kr.com",
		HostTitle:       "36kr",
		HostDescription: "36氪_让一部分人先看到未来",

		TargetURl: "https://36kr.com/newsflashes",

		ListContainers: ".flow-item",
		Title:          ".item-main .newsflash-item a",
		URL:            ".item-main .newsflash-item a",
		Date:           ".item-related span.time",
		Contents:       ".item-desc span",
		Author:         ".item-related a.project span",
		Description:    ".article-item-description",
	}

	toRes := rsshub.ToResponseRod{}
	doc, _ := toRes.Response(rule.TargetURl)

	feeds := rsshub.NewFeeds()
	f, _ := feeds.ToFeed(doc, rule)
	fmt.Println(f.ToRss())

}

func NewDaYu() {
	rootURL := rsshub.Rule{
		TargetURl: "https://btcdayu.gitbook.io/dayu",
		URL:       "div.css-175oi2r a",
	}
	rule := rsshub.Rule{
		Host:            "https://btcdayu.gitbook.io/dayu",
		HostTitle:       "聪明的投资者(币圈版)",
		HostDescription: "聪明的投资者(币圈版)",

		TargetURl: "https://btcdayu.gitbook.io/dayu/cong-ming-de-tou-zi-zhe/di-yi-zhang-tou-zi-yu-tou-ji/1.-cong-ming-tou-zi-zhe-de-san-da-yuan-ze",

		ListContainers: "main",
		Title:          "div.css-175oi2r h1",
		URL:            ".item-main .newsflash-item a",
		Date:           ".item-related span.time",
		Contents:       ".css-175oi2r.r-bnwqim",
		Author:         ".item-related a.project span",
		Description:    ".css-175oi2r div['dir'='auto']",
	}

	toRes := rsshub.ToResponseRod{}

	rootDoc, _ := toRes.Response(rootURL.TargetURl)
	var targetUrls []string
	rootDoc.Find(rootURL.URL).Each(func(i int, selection *goquery.Selection) {
		v, _ := selection.Attr("href")
		if strings.Contains(v, "/dayu/") && len(v) > 10 {
			targetUrls = append(targetUrls, fmt.Sprintf("%s%s", rootURL.TargetURl, v[5:]))
		}
	})

	toRes.Response("https://btcdayu.gitbook.io/dayu/cong-ming-de-tou-zi-zhe/di-yi-zhang-tou-zi-yu-tou-ji")

	for _, one := range targetUrls {
		fmt.Println(one)
		doc, err := toRes.Response(one)
		if err != nil {
			continue
		}
		feeds := rsshub.NewFeeds()
		f, _ := feeds.ToFeed(doc, rule)
		if len(f.Items) > 0 {
			fmt.Println(f.ToRss())
		}
	}

}
