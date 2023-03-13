package schema

import (
	"database/sql"
	"reflect"
)

type Users struct {
	ID          uint         `gorm:"primarykey"`
	CreateTime  sql.NullTime `gorm:"column:create_time;type:datetime;autoCreateTime" json:"create_time"`
	UpdateTime  sql.NullTime `gorm:"column:update_time;type:datetime;autoCreateTime" json:"update_time"`
	Uid         string       `gorm:"column:uid;type:varchar(32);comment:用户" json:"uid"`
	Description string       `gorm:"column:description;type:varchar(128);comment:简介" json:"description"`
}

func (u Users) TableName() string {
	return "users"
}

func (u Users) Fields() map[string]string {
	t := reflect.TypeOf(u)
	var m = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		m[t.Field(i).Name] = t.Field(i).Tag.Get("json")
	}
	return m
}
