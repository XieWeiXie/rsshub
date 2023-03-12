package schema

import (
	"database/sql"
	"reflect"
)

type Unique struct {
	ID         uint         `gorm:"primarykey"`
	CreateTime sql.NullTime `gorm:"column:create_time;type:datetime;autoCreateTime" json:"create_time"`
	UpdateTime sql.NullTime `gorm:"column:update_time;type:datetime;autoCreateTime" json:"update_time"`
	Uid        string       `gorm:"column:uid;type:varchar(32);comment:用户标识" json:"uid"`
	Url        string       `gorm:"column:url;type:varchar(64);comment:路由" json:"url"`
	Type       string       `gorm:"column:type;type:varchar(12);comment:来源区分" json:"type"`
}

func (u Unique) TableName() string {
	return "unique_url"
}

func (u Unique) Fields() map[string]string {
	t := reflect.TypeOf(u)
	var m = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		m[t.Field(i).Name] = t.Field(i).Tag.Get("json")
	}
	return m
}
