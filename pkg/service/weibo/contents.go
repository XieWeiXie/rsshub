package weibo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/XieWeiXie/rsshub/pkg/db"
	"github.com/XieWeiXie/rsshub/pkg/hub"
	"github.com/XieWeiXie/rsshub/pkg/schema"
	"github.com/gorilla/feeds"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"strings"
)

type Contents interface {
	ToContent(url string) error
	ForContent(urls string) (map[string][]byte, error)
}

func NewContents() Contents {
	return contents{}
}

type contents struct {
}

type contentRule struct {
	List    string
	Content string
	Images  string
	Date    string
	Title   string
	Uid     string
}

func (c contents) ToContent(url string) error {
	toResponse := hub.NewToResponse(http.Header{
		"Cookie":     []string{os.Getenv("Cookie")},
		"User-Agent": []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"},
	})
	doc, err := toResponse.Response(url)
	if err != nil {
		return err
	}

	if strings.Contains(doc.Find("div.me").Text(), "Permission Denied!") || strings.Contains(doc.Find("div.me").Text(), "Due to author settings, The contents cannot be viewed.") {
		return errors.New("Permission Denied!")
	}

	var rule = contentRule{
		List:    "#M_",
		Content: "div span.ctt",
		Images:  "div a", // first
		Date:    "div span.ct",
		Title:   "div a", // last
		Uid:     "div a",
	}

	var feed = new(feeds.Feed)
	var items = new(feeds.Item)
	items.Link = &feeds.Link{
		Href: fmt.Sprintf("%s/%s/%s", pc, toUrlUid(url), toUrlValue(url)),
		Rel:  url,
	}
	feed.Link = items.Link
	var uid = toUrlUid(url)
	doc.Find(rule.List).Each(func(i int, selection *goquery.Selection) {
		title := toStringReg(selection.Find(rule.Title).Eq(0).Text())
		feed.Title = title
		feed.Author = &feeds.Author{
			Name: title,
		}
		items.Title = title
		items.Author = feed.Author

		c := selection.Find(rule.Content).Text()
		c = toStringReplace(c)
		items.Description = c
		if len([]rune(c)) > 10 {
			items.Description = string([]rune(c)[:10]) + "..."
		}
		feed.Description = items.Description
		items.Content = c
		if selection.Find("div").Length() > 1 {
			image, ok := selection.Find("div").Eq(1).Find(rule.Images).First().Attr(href)
			if ok {
				image = fmt.Sprintf("%s%s", rootNoSlash, image)
				feed.Image = &feeds.Image{
					Url:  image,
					Link: image,
				}
			}
		}
		date := toStringReg(selection.Find(rule.Date).Text())
		items.Created = toDate(date)

	})
	feed.Add(items)
	by, _ := json.Marshal(feed)
	if len(by) == 0 {
		return nil
	}
	op := schema.ContentsOp{DB: *db.DefaultMysql}
	req := schema.RequestContentsOp{}
	return op.ToOne(
		schema.NewRequestContentOp(
			req.WithRequestContentOpContent(map[string]string{feed.Link.Href: string(by)}),
			req.WithRequestContentOpUID(uid),
			req.WithRequestContentOpPublish(items.Created),
		),
	)

}

func (c contents) ForContent(url string) (map[string][]byte, error) {
	op := schema.ContentsOp{}
	content, err := op.ForOne(
		url,
	)
	if err != nil {
		return nil, err
	}
	var m = make(map[string][]byte)
	var b bytes.Buffer
	b.WriteString(content.Content)
	m[url] = b.Bytes()
	return m, nil
}
