package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpcProgram/server/service"
	"log"
)

/**
* @Author:liyao
* @Description:
* @Version:
* @File:grpc-client
* @Date: 2022/11/28-17:50
 */
func main() {
	ctx := context.Background()
	conn, err := grpc.Dial(":8001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("无法连接服务端", err)
	}
	defer conn.Close()
	productClient := service.NewProductServiceClient(conn)
	productReq := &service.ProductReq{
		ProductId: 1000,
	}
	productStockResp, err := productClient.GetProductStock(ctx, productReq)
	if err != nil {
		log.Fatalln("从服务端查询库存失败", err)
	}
	log.Print("从服务端查询库存成功,库存为", productStockResp.ProductStock)

}
