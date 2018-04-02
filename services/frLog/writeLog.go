package frLog

import (
	"strings"
	"encoding/json"
	"time"
	"net"
	"frrpc/services"
	pb "frrpc/protoFile"
	"frrpc/const"
	"log"
	"github.com/garyburd/redigo/redis"
)

/**
 * 获取redis链接句柄 写日志比较特殊  专用redis服务器 故单独写在此处
 */
func GetRedisClient() redis.Conn {
	var client redis.Conn
	var host , pass string
	host = "120.76.201.199:6479"
	pass = "VWCwmQKwjYLJWB71ccT1"
	client , err :=  redis.Dial("tcp", host)
	if err != nil {
		log.Fatal("redis error: %v", err)
	}
	_ , err = client.Do("AUTH", pass)
	if err != nil {
		log.Fatal("redis error: %v", err)
	}
	return client
}

/**
 * 写入kibana缓存 生成日志
 * @param tag string 日志标签
 * @param info string 日志信息
 * @param level string 日志级别
 */
func WriteLog(tag string, info string, level string) (*pb.FrLogReply, error) {
	var prefix string
	baseServ := services.BaseService{}
	log_data := make(map[string]interface{})
	if baseServ.GetEnv() == "prod" {
		prefix = ""
	} else {
		prefix = "test_"
	}
	log_data = __getBaseRecord(level,info)
	log_data["tags"] = strings.ToLower(prefix + tag)
	log_data["type"] = strings.ToLower(prefix + tag)
	log_json , _ := json.Marshal(log_data)

	client := GetRedisClient()
	_ , err := client.Do("LPUSH","common_api_access_log", string(log_json))
	defer client.Close()
	reply := pb.FrLogReply{
		Code:_const.STATUS_SUCCESS,
		Message:"success",
		Data:nil,
	}
	if err != nil {
		reply.Code = _const.LOG_WRITE_ERR
		reply.Message = "写入日志出错"
		reply.Data = nil
	}
	baseServ.LogInfo("frlog_request" , "tag:"+tag+" info:"+info+" level:"+level+" status:"+reply.Message)
	return &reply, err
}

/**
 * 构造写入日志内容
 * @param level string 日志级别
 * @param info string 日志信息
 * @return map[string]interface{}
 */
func __getBaseRecord(level string,info string) map[string]interface{} {
	log_data := make(map[string]interface{})
	log_data["log_time"] = time.Now().Unix()
	log_data["level"]    = level
	log_data["server"],_ = net.InterfaceAddrs()
	log_data["msg"]      = info
	return log_data
}