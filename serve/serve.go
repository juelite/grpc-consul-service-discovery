package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "frrpc/protoFile"
	"frrpc/services/redisCache"
	"frrpc/services/frLog"
	"frrpc/services"
	"flag"
	"time"
	"frrpc/consul"
)



type server struct {}

var (
	serv = flag.String("service", "Cache", "elcLog RedisCache")
	port = flag.Int("port", 50051, "listening port")
	reg  = flag.String("reg", "114.55.248.175:8500", "register address")
	ttl = 15
)


/**
 * 写入redis缓存服务
 */
func (s *server) RedisCache(ctx context.Context, in *pb.RedisCacheRequest) (reply *pb.RedisCacheReply , err error) {
	reply , err = redisCache.RedisCache(in.Name , in.Express , in.Value)
	return
}

/**
 * 获取redis缓存
 */
func (s *server) GetCache(ctx context.Context, in *pb.GetCacheRequest) (reply *pb.GetCacheReply , err error) {
	reply , err = redisCache.GetCache(in.Name)
	return
}

/**
 * 写kibana日志服务
 */
func (s *server) FrLog(ctx context.Context , in *pb.FrLogRequest) (reply *pb.FrLogReply , err error) {
	reply , err = frLog.WriteLog(in.Tag , in.Info , in.Level)
	return
}

func main() {
	flag.Parse()
	base := services.BaseService{}
	//起RPC服务
	gRPCport := base.GetVal("rpcserve")
	base.LogInfo("index" , gRPCport)
	lis , err := net.Listen("tcp", gRPCport)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	//注册服务到consul
	err = consul.Register(*serv, "114.55.248.175", *port, *reg, time.Second*10, ttl)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}
