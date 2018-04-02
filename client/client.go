package main

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "frrpc/protoFile"
	"fmt"
)

const (
	address     = "114.55.248.175:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	//写缓存测试
	request := &pb.RedisCacheRequest{
		Name:"wangyu",
		Express:2000,
		Value:"hello",
	}
	r , err := c.RedisCache(context.Background(), request)
	/*request := &pb.GetCacheRequest{
		Name:"wangyu",
	}
	r , err := c.GetCache(context.Background(), request)*/
	/*
	写日志测试
	request := &pb.FrLogRequest{
		Tag:"ceshirizhi",
		Info:"hello",
		Level:"info",
	}
	r, err := c.FrLog(context.Background(), request)
	*/
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	fmt.Println(r.Code , r.Message , r.Data)
}