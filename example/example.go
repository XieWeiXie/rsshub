package example

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/XieWeiXie/rsshub/pkg/hub"
	"strings"
)

func New36krNews() {
	rule := hub.Rule{
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

	toRes := hub.ToResponseRod{}
	doc, _ := toRes.Response(rule.TargetURl)

	feeds := hub.NewFeeds()
	f, _ := feeds.ToFeed(doc, rule)
	fmt.Println(f.ToRss())

}

func NewDaYu() {
	rootURL := hub.Rule{
		TargetURl: "https://btcdayu.gitbook.io/dayu",
		URL:       "div.css-175oi2r a",
	}
	rule := hub.Rule{
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

	toRes := hub.ToResponseRod{}

	rootDoc, _ := toRes.Response(rootURL.TargetURl)
	var targetUrls []string
	rootDoc.Find(rootURL.URL).Each(func(i int, selection *goquery.Selection) {
		v, _ := selection.Attr("href")
		if strings.Contains(v, "/dayu/") && len(v) > 10 {
			targetUrls = append(targetUrls, fmt.Sprintf("%s%s", rootURL.TargetURl, v[5:]))
		}
	})

	for _, one := range targetUrls {
		doc, err := toRes.Response(one)
		if err != nil {
			continue
		}
		feeds := hub.NewFeeds()
		f, _ := feeds.ToFeed(doc, rule)
		if len(f.Items) > 0 {
			fmt.Println(f.ToRss())
		}
	}

}
