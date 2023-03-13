package controller

import (
	"context"
	"fmt"
	v1 "github.com/XieWeiXie/rsshub/api/v1"
	"github.com/XieWeiXie/rsshub/pkg/grpc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WeiboCnAdd struct {
}

func (w WeiboCnAdd) Describe() string {
	describe := "WeiBoAdd"
	cacheHandler[describe] = w
	return describe
}

func (w WeiboCnAdd) ToRSSHandler(ctx *gin.Context) {
	var req = new(v1.NewAddWeiboReq)
	if ctx.ShouldBind(&req) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": "参数错误",
		})
		return
	}
	client := v1.NewWeiboClient(grpcservice.GrpcClient(9091))
	_, err := client.NewAddWeibo(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":    "服务端错误",
			"message": fmt.Sprintf("err: #%v", err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, "OK")
}
