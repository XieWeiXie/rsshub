package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DefaultMysql *gorm.DB
)

const (
	WeiboDatabase = "c_weibo"
)

func Mysql(database string) {
	db, err := gorm.Open(
		mysql.New(mysql.Config{
			DriverName:                "mysql",
			ServerVersion:             "8.0.32",
			DSN:                       fmt.Sprintf("root:admin123456@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", database),
			DSNConfig:                 nil,
			Conn:                      nil,
			SkipInitializeWithVersion: false,
			DefaultStringSize:         256,
		}),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		return
	}
	DefaultMysql = db
}
