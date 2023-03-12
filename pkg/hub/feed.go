package hub

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/feeds"
	"strings"
	"time"
)

type Feeds interface {
	ToFeed(document *goquery.Document, rules Rule) (*feeds.Feed, error)
}

func NewFeeds() Feeds {
	return toFeed{}
}

type toFeed struct {
}

const href = "href"

func (t toFeed) ToFeed(document *goquery.Document, rules Rule) (fs *feeds.Feed, err error) {
	toTrimSpace := func(s string) string {
		return strings.TrimSpace(s)
	}
	fs = new(feeds.Feed)
	fs.Title = rules.HostTitle
	fs.Link = &feeds.Link{Href: rules.Host}
	fs.Description = rules.HostDescription
	document.Find(rules.ListContainers).Each(func(i int, selection *goquery.Selection) {
		var one = new(feeds.Item)
		one.Title = toTrimSpace(selection.Find(rules.Title).Text())
		v, _ := selection.Find(rules.URL).Attr(href)
		one.Link = &feeds.Link{Href: fmt.Sprintf("%s%s", rules.Host, toTrimSpace(v))}
		one.Author = &feeds.Author{Name: toTrimSpace(selection.Find(rules.Author).Text())}
		one.Description = toTrimSpace(selection.Find(rules.Description).Text())
		one.Created = time.Now()
		one.Content = toTrimSpace(selection.Find(rules.Contents).Text())
		fs.Add(one)
	})
	return
}
