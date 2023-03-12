package schema

import (
	"database/sql"
	"reflect"
)

type Contents struct {
	ID          uint         `gorm:"primarykey"`
	CreateTime  sql.NullTime `gorm:"column:create_time;type:datetime;autoCreateTime" json:"create_time"`
	UpdateTime  sql.NullTime `gorm:"column:update_time;type:datetime;autoCreateTime" json:"update_time"`
	Url         string       `gorm:"column:url;type:varchar(32);comment:原始链接" json:"url"`
	Content     string       `gorm:"column:content;type:text;comment:内容 XML" json:"content"`
	Uid         string       `gorm:"column:uid;type:varchar(32);comment:用户" json:"uid"`
	PublishTime sql.NullTime `gorm:"column:publish_time;type:datetime;comment:发布时间" json:"publish_time"`
}

func (c Contents) TableName() string {
	return "contents"
}

func (c Contents) Fields() map[string]string {
	t := reflect.TypeOf(c)
	var m = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		m[t.Field(i).Name] = t.Field(i).Tag.Get("json")
	}
	return m
}
