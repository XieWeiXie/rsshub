package hub

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"time"
)

type Pager interface {
	ToPage(url string) (string, error)
}

func NewPage(_type int, values map[string]string) Pager {
	switch _type {
	case 1:
		return normalPager{}
	case 2:
		return loginPager{values: values}
	}
	return nil
}

type normalPager struct {
}

func (n normalPager) ToPage(url string) (string, error) {
	var page *rod.Page
	l := launcher.MustResolveURL(defaultProxyUrl)
	page = rod.New().ControlURL(l).MustConnect().MustPage(url)
	page.
		Timeout(5 * time.Second).
		MustWaitLoad().
		MustElement("title").
		CancelTimeout().
		Timeout(10 * time.Second).
		MustText()
	defer page.MustClose()
	return page.HTML()
}

type loginPager struct {
	values map[string]string
}

func (l loginPager) ToPage(url string) (string, error) {
	var page *rod.Page
	var cookies []*proto.NetworkCookieParam
	for k, v := range l.values {
		cookies = append(cookies, &proto.NetworkCookieParam{
			Name:   k,
			Value:  v,
			Domain: url,
		})
	}
	browser := rod.New().ControlURL(launcher.MustResolveURL(defaultProxyUrl)).MustConnect()
	if len(cookies) > 0 {
		if err := browser.SetCookies(cookies); err != nil {
			return "", err
		}
	}
	page = browser.MustPage(url)
	page.
		Timeout(5 * time.Second).
		MustWaitLoad().
		MustElement("title").
		CancelTimeout().
		Timeout(10 * time.Second).
		MustText()
	defer page.MustClose()
	return page.HTML()
}
