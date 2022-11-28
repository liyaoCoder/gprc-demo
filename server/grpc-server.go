package main

import (
	"google.golang.org/grpc"
	"grpcProgram/server/service"
	"log"
	"net"
)

/**
* @Author:liyao
* @Description:
* @Version:
* @File:grpc-server
* @Date: 2022/11/28-17:38
 */
func main() {
	log.Print("开始启动grpc-server服务")
	rpcServer := grpc.NewServer()
	service.RegisterProductServiceServer(rpcServer, &service.Product)
	listener, _ := net.Listen("tcp", ":8001")
	err := rpcServer.Serve(listener)
	if err != nil {
		log.Fatalln("启动服务出错", err)
	}
	log.Print("启动服务成功")

}
