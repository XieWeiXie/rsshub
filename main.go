package main

import (
	"fmt"
	"github.com/XieWeiXie/rsshub/controller"
	"github.com/XieWeiXie/rsshub/pkg/db"
	grpcservice "github.com/XieWeiXie/rsshub/pkg/grpc"
	"github.com/XieWeiXie/rsshub/pkg/task"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	engine.Use(gin.Recovery(), gin.Logger())
	db.Mysql(db.WeiboDatabase)
	go task.HandlerTask()

	v1 := engine.Group("rssHub/v1")

	v1.GET("/hi", controller.NewHandler(fmt.Sprintf("%s", controller.Hello{}.Describe())).ToRSSHandler)
	v1.GET("/36kr", controller.NewHandler(fmt.Sprintf("%s", controller.Kr36{}.Describe())).ToRSSHandler)
	v1.GET("/weibo/:user", controller.NewHandler(fmt.Sprintf("%s", controller.WeiboCn{}.Describe())).ToRSSHandler)

	go grpcservice.GrpcService(9091)

	engine.Run(":8080")
}
