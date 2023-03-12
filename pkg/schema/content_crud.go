package schema

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type ContentsOp struct {
	gorm.DB
}

type RequestContentsOp struct {
	content map[string]string
	uid     string
	publish time.Time
}

func (c RequestContentsOp) WithRequestContentOpContent(content map[string]string) func(r *RequestContentsOp) {
	return func(r *RequestContentsOp) {
		r.content = content
	}
}

func (c RequestContentsOp) WithRequestContentOpUID(uid string) func(r *RequestContentsOp) {
	return func(r *RequestContentsOp) {
		r.uid = uid
	}
}

func (c RequestContentsOp) WithRequestContentOpPublish(pub time.Time) func(r *RequestContentsOp) {
	return func(r *RequestContentsOp) {
		r.publish = pub
	}
}

func NewRequestContentOp(funcs ...func(r *RequestContentsOp)) RequestContentsOp {
	var req = new(RequestContentsOp)
	for _, f := range funcs {
		f(req)
	}
	return *req
}

func (c ContentsOp) ToOne(req RequestContentsOp) error {
	if len(req.content) != 1 {
		return nil
	}
	var (
		key   string
		value string
	)

	for k, v := range req.content {
		key = k
		value = v
		break
	}
	return c.Model(Contents{}).Create(&Contents{
		Url:     key,
		Content: value,
		Uid:     req.uid,
		PublishTime: sql.NullTime{
			Time:  req.publish,
			Valid: true,
		},
	}).Error
}

func (c ContentsOp) ForOne(req string) (Contents, error) {
	var doc Contents
	err := c.Model(Contents{}).Where(fmt.Sprintf("%s = ?", Contents{}.Fields()["Url"]), req).First(&doc).Error
	return doc, err
}

func (c ContentsOp) CanFull(req string) (Contents, error) {
	var doc Contents
	err := c.Model(Contents{}).Where(fmt.Sprintf("%s = ?", Contents{}.Fields()["Uid"]), req).First(&doc).Error
	return doc, err
}

func (c ContentsOp) ToContents(req RequestContentsOp) error {
	var docs []Contents
	for k, v := range req.content {
		docs = append(docs, Contents{
			Url:     k,
			Content: v,
			Uid:     req.uid,
		})
	}
	return c.Model(Contents{}).CreateInBatches(docs, 1000).Error
}

type RequestContentQuery struct {
	selectors []string
	uid       string
}

func (r *RequestContentQuery) WithRequestContentQuery(uid string) func(r *RequestContentQuery) {
	return func(r *RequestContentQuery) {
		r.uid = uid
	}
}

func (r *RequestContentQuery) WithRequestContentSelector(selectors []string) func(r *RequestContentQuery) {
	return func(r *RequestContentQuery) {
		r.selectors = selectors
	}
}

func (c ContentsOp) Batch(req RequestContentQuery) *gorm.DB {
	return c.Model(Contents{}).Where(fmt.Sprintf("%s = ?", Contents{}.Fields()["Uid"]), req.uid).
		Order(
			clause.OrderByColumn{
				Column: clause.Column{
					Table: clause.CurrentTable, Name: Contents{}.Fields()["PublishTime"],
				},
				Desc:    true,
				Reorder: false,
			},
		).
		Limit(-1)
}
