package task

import (
	"github.com/XieWeiXie/rsshub/pkg/service/weibo"
	"github.com/robfig/cron"
)

func HandlerTask() {
	c := cron.New()
	_ = c.AddJob("0 */3 * * * ?", NewWeiBo{})
	c.Start()
}

type NewWeiBo struct {
}

func (n NewWeiBo) Run() {
	s := weibo.Service{}
	s.ContentsFull("1667517061")
}
