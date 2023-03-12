package task

import (
	"github.com/XieWeiXie/rsshub/pkg/db"
	"github.com/XieWeiXie/rsshub/pkg/schema"
	"github.com/XieWeiXie/rsshub/pkg/service/weibo"
	"github.com/robfig/cron"
	"gorm.io/gorm"
)

func HandlerTask() {
	c := cron.New()
	_ = c.AddJob("0 */10 * * * ?", NewWeiboURL{})
	_ = c.AddJob("20 */30 * * * ?", NewWeiBoContent{})
	c.Start()
}

type NewWeiBoContent struct {
}

func (n NewWeiBoContent) Run() {
	s := weibo.Service{}
	for k, _ := range can() {
		_ = s.ContentsFull(k)
	}
}

func can() map[string]struct{} {
	op := schema.UserOp{DB: *db.DefaultMysql}
	var (
		users  []schema.Users
		unique = make(map[string]struct{})
	)
	op.ForAll().FindInBatches(&users, 10, func(tx *gorm.DB, batch int) error {
		for _, one := range users {
			unique[one.Uid] = struct{}{}
		}
		return nil
	})

	op2 := schema.UniqueOp{DB: *db.DefaultMysql}
	var can = make(map[string]struct{})
	for k, _ := range unique {
		v, _ := op2.CanFull(k)
		if v {
			can[k] = struct{}{}
		}
	}
	if len(can) == 0 {
		return nil
	}
	return can
}

type NewWeiboURL struct {
}

func (n NewWeiboURL) Run() {
	can := can()
	s := weibo.Service{}
	for k, _ := range can {
		_ = s.FistFull(k)
	}
}
