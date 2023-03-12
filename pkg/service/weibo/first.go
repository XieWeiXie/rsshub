package weibo

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/XieWeiXie/rsshub/pkg/hub"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type First interface {
	ToFirst(func([]string)) error
}

func NewFirst(uid string, page int) First {
	return firstInit{uid: uid, page: page}
}

const (
	root        = "https://weibo.cn/"
	rootNoSlash = "https://weibo.cn"
	href        = "href"
	value       = "value"
	filter      = "filter"
	page        = "page"
	pc          = "https://weibo.com"
)

type firstInit struct {
	uid  string
	page int
}

func (f firstInit) RootURL() string {
	return fmt.Sprintf("%s%s?filter=1", root, f.uid)
}

func (f firstInit) Next(page int) string {
	return fmt.Sprintf("%s%s?filter=1&page=%d", root, f.uid, page)
}

type urlRule struct {
	List  string
	Url   string
	Page  string
	Value string
}

func (f firstInit) ToFirst(do func([]string)) (err error) {
	toResponse := hub.NewToResponse(
		http.Header{
			"Cookie":     []string{os.Getenv("Cookie")},
			"User-Agent": []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"},
		})
	doc, err := toResponse.Response(f.RootURL())
	if err != nil {
		log.Fatalln("get response fail")
		return err
	}

	var rl = urlRule{
		List:  "div.c div a.cc",
		Page:  "#pagelist form div",
		Value: "value",
	}
	var (
		page string = strconv.Itoa(f.page)
		pa          = 2
	)

	if _, ok := doc.Find("body").Attr("div"); !ok {
		log.Println("更新Cookie")
		return
	}
	if f.page == 0 {
		pa = 1
		pageList := doc.Find(rl.Page).First()
		for _, one := range pageList.Children().Get(1).Attr {
			if one.Key == value {
				page = one.Val
				break
			}
		}
	}

	var pageURL []string
	toPageURL := func(doc *goquery.Document, reg string) []string {
		var everyURL []string
		doc.Find(reg).Each(func(i int, selection *goquery.Selection) {
			pageUrl, ok := selection.Attr(href)
			if ok {
				everyURL = append(everyURL, pageUrl)
			}
		})
		return everyURL
	}

	l := toPageURL(doc, rl.List)
	if len(l) > 0 {
		do(l)
	}

	pageURL = append(pageURL, l...)
	pageInt, _ := strconv.Atoi(page)
	for pa <= pageInt {
		doc, err = toResponse.Response(f.Next(pa))
		if err != nil {
			continue
		}
		l := toPageURL(doc, rl.List)
		if len(l) > 0 {
			do(l)
		}
		if len(l) == 0 {
			break
		}
		pageURL = append(pageURL, l...)
		pa += 1
		time.Sleep(2 * time.Second)
	}
	for _, one := range pageURL {
		log.Println(fmt.Sprintf("uid: %s url: %s", f.uid, one))
	}
	return
}
