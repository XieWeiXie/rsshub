package controller

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	v1 "github.com/XieWeiXie/rsshub/api/v1"
	grpcservice "github.com/XieWeiXie/rsshub/pkg/grpc"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/feeds"
	"io"
	"log"
	"net/http"
)

type WeiboCn struct {
}

func (w WeiboCn) Describe() string {
	describe := "WeiBo"
	cacheHandler[describe] = w
	return describe
}

func (w WeiboCn) ToRSSHandler(ctx *gin.Context) {
	client := v1.NewWeiboClient(grpcservice.GrpcClient(9091))
	stream, err := client.NewAllWeibo(context.Background(), &v1.NewAllWeiboReq{Uid: ctx.Param("user")})
	if err != nil {
		ctx.Data(http.StatusBadRequest, defaultXML, nil)
		return
	}
	var m = make(map[string][]byte)
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		if len(res.Response) > 0 {
			for k, v := range res.Response {
				m[k] = v
			}
		}
	}
	type Feed struct {
		XMLName xml.Name `xml:"rss"`
		Version string   `xml:"version,attr"`
		Channel feeds.Item
	}
	var feed = new(feeds.Feed)
	for _, v := range m {
		var oneFeed feeds.Feed
		if err = json.Unmarshal(v, &oneFeed); err != nil {
			log.Println("xml unmarshal fail, err", err)
			break
		}
		feed.Items = append(feed.Items, oneFeed.Items...)
	}
	oneItem := feed.Items[0]
	feed.Title = oneItem.Title
	feed.Author = oneItem.Author
	feed.Description = "微博精选"
	feed.Link = &feeds.Link{
		Href: fmt.Sprintf("https://weibo.cn/%s", ctx.Param("user")),
		Rel:  fmt.Sprintf("https://weibo.cn/%s", ctx.Param("user")),
	}
	b, _ := feed.ToRss()
	ctx.Data(http.StatusOK, defaultXML, []byte(b))

}
