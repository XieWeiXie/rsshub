package hub

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

const defaultUserAgentValue = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"
const defaultUserAgentKey = "User-Agent"

type ToResponse interface {
	Response(url string) (*goquery.Document, error)
}

func NewToResponse(headers http.Header) ToResponse {
	return &ToResponseHTTP{headers: headers}
}
func NewToResponseV2(pager Pager) ToResponse {
	return &ToResponseRod{pager}
}

type ToResponseHTTP struct {
	headers http.Header
}

func (t *ToResponseHTTP) Response(url string) (*goquery.Document, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if t.headers != nil {
		request.Header = t.headers
	}
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
	Pager
}

var defaultProxyUrl = "127.0.0.1:9222"

func (t *ToResponseRod) Response(url string) (*goquery.Document, error) {
	str, err := t.Pager.ToPage(url)
	if err != nil {
		return nil, err
	}
	return goquery.NewDocumentFromReader(strings.NewReader(str))
}
