package controller

import (
	"github.com/XieWeiXie/rsshub/pkg/hub"
	"testing"
)

func TestWeiboCn(t *testing.T) {
	rule := &hub.Rule{
		Host:            "https://weibo.cn/",
		HostTitle:       "微博",
		HostDescription: "微博",
	}
	rule.ToTarget("https://weibo.cn/comment/MwFyL2GkK?ckAll=1").
		ToContainers("#M_").
		ToContents("span.ctt")

	toResponse := hub.NewToResponse(nil)
	toResponse.Response(rule.TargetURl)

}
