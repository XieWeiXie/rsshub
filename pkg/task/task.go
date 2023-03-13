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
	for k, _ := range can(1) {
		_ = s.ContentsFull(k)
	}
}

func can(_type int) map[string]struct{} {
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
	var can = make(map[string]struct{})

	switch _type {
	case 1:
		op2 := schema.UniqueOp{DB: *db.DefaultMysql}
		for k, _ := range unique {
			v, _ := op2.CanFull(k)
			if v {
				can[k] = struct{}{}
			}
		}
	case 2:
		op2 := schema.ContentsOp{DB: *db.DefaultMysql}
		for k, _ := range unique {
			v, _ := op2.CanFull(k)
			if v.ID == 0 {
				can[k] = struct{}{}
			}
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
	can := can(2)
	s := weibo.Service{}
	for k, _ := range can {
		_ = s.FistFull(k)
	}
}
