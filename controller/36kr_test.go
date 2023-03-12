package controller

import (
	"fmt"
	"github.com/XieWeiXie/rsshub/pkg/hub"
	"testing"
)

func Test36Kr(t *testing.T) {
	toRes := hub.ToResponseRod{}
	doc, _ := toRes.Response("https://36kr.com/")
	fmt.Println(doc)
}
