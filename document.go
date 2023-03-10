package rsshub

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"net/http"
	"strings"
	"time"
)

const defaultUserAgentValue = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"
const defaultUserAgentKey = "User-Agent"

type ToResponse interface {
	Response(url string) (*goquery.Document, error)
}

type ToResponseHTTP struct {
	url string
}

func WithToResponseHTTPURL(url string) func(r *ToResponseHTTP) {
	return func(r *ToResponseHTTP) {
		r.url = url
	}
}

func (t *ToResponseHTTP) Response(url string) (*goquery.Document, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add(defaultUserAgentKey, defaultUserAgentValue)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	return goquery.NewDocumentFromReader(res.Body)
}

type ToResponseRod struct {
}

var defaultProxyUrl = "127.0.0.1:9222"

func (t *ToResponseRod) Response(url string) (*goquery.Document, error) {
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
	str, err := page.HTML()
	if err != nil {
		return nil, err
	}
	return goquery.NewDocumentFromReader(strings.NewReader(str))
}
