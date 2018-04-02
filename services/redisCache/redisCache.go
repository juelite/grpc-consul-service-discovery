package redisCache

import (
	"frrpc/services"
	"frrpc/const"
	pb "frrpc/protoFile"
)


/**
 * 写入redis缓存
 * @param name string 键名
 * @param exp int64 有效期，0为不过期，单位秒
 * @param val string 值
 * @return err error 写入结果
 */
func RedisCache(name string , exp int64 , val string) (*pb.RedisCacheReply , error) {
	base := services.BaseService{}
	base.LogInfo("redis_cache_request" , "name:"+name+" val:"+val)
	var err error
	client := base.GetRedisClient()
	if exp > 0 {
		_ , err = client.Do("SETEX" , name , exp , val)
	} else {
		_ , err = client.Do("SET" , name , val)
	}

	ret := pb.RedisCacheReply{
		Code:_const.STATUS_SUCCESS,
		Message:"success",
		Data:nil,
	}
	if err != nil {
		ret.Code = _const.REDIS_WRITE_ERR
		ret.Message = "缓存写入失败"
	}
	return &ret , nil
}

func GetCache(name string) (*pb.GetCacheReply , error) {
	base := services.BaseService{}
	var (
		err error
		data string
	)
	client := base.GetRedisClient()
	res , err := client.Do("GET" , name)
	if res != nil {
		data = string(res.([]byte))
	}
	rsp := make(map[string]string)
	ret := pb.GetCacheReply{
		Code:_const.STATUS_SUCCESS,
		Message:"success",
	}
	if err != nil {
		ret.Code = _const.REDIS_READ_ERR
		ret.Message = "读取缓存失败"
	}
	if data != "" {
		rsp[name] = data
	}
	ret.Data = rsp
	return &ret , nil
}