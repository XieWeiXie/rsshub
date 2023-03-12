package grpcservice

import (
	"fmt"
	v1 "github.com/XieWeiXie/rsshub/api/v1"
	"github.com/XieWeiXie/rsshub/pkg/service/weibo"
	"google.golang.org/grpc"
	"net"
)

func GrpcService(port int) {
	listen, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		fmt.Println(err)
	}
	s := grpc.NewServer(
		grpc.MaxRecvMsgSize(8*1024*1024),
		grpc.MaxSendMsgSize(8*1024*1024),
	)
	v1.RegisterWeiboServer(s, &weibo.Service{})
	s.Serve(listen)
}

func GrpcClient(port int) grpc.ClientConnInterface {
	conn, err := grpc.Dial(fmt.Sprintf(":%d", port), grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return conn
}
