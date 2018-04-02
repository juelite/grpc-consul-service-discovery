package services

import (
	"github.com/Unknwon/goconfig"
	"os"
	"time"
	"log"
	"io"
	"github.com/garyburd/redigo/redis"
)

type BaseService struct {

}

//相对路径是根据可执行文件当前位置
const confPath  =  "./conf/setting.conf"
const envPath  =  "./conf/env.conf"

var env string

/**
 * 获取环境变量
 */
func (t *BaseService) GetEnv() string {
	c , err := goconfig.LoadConfigFile(envPath)
	if err != nil {
		log.Fatal("load config file error: ", err)
	}
	//先获取运行环境
	env , err = c.GetValue("default" , "runmode")
	if err != nil {
		log.Fatal("load env file error: ", err)
	}
	return env
}

/**
 * 获取配置文件
 */
func (t *BaseService) GetConf() (*goconfig.ConfigFile , error) {
	c , err := goconfig.LoadConfigFile(confPath)
	return c , err
}

/**
 * 根据键获取值
 */
func (t *BaseService) GetVal(name string) string {
	cfg , err := t.GetConf()
	if err != nil {
		log.Fatal("load config file error: ", err)
	}
	section := t.GetEnv()
	val , err := cfg.GetValue(section , name)
	if err != nil {
		log.Fatal("parse config fail 2: ", err)
	}
	return val
}

/**
 * 日志写入，需先在调用的模块controller统计目录创建runtime目录
 * @param file_name string 要写入的文件
 * @param file_content string 要写入的内容
 */
func (t *BaseService) LogInfo(file_name string, file_content string) {
	//文件夹路径
	file, err := os.Open("runtime/" + time.Now().Format("2006-01-02"))
	if err != nil {
		os.Mkdir("runtime/"+time.Now().Format("2006-01-02"), os.ModePerm)
	}
	//确定日志路径
	file_name = "runtime/" + time.Now().Format("2006-01-02") + "/" + file_name
	//打开文件,不存在则创建
	file, err = os.OpenFile(file_name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		println(err.Error())
	}
	//确定输出格式
	trace := log.New(io.MultiWriter(file, os.Stdin), "Info:", log.Ldate|log.Ltime)
	//写入日志内容
	trace.Println(file_content)
}

/**
 * 获取redis链接
 */
func (t *BaseService) GetRedisClient() redis.Conn{
	var client redis.Conn
	var host , pass string
	host = t.GetVal("redishost")
	pass = t.GetVal("redispass")
	client , err :=  redis.Dial("tcp", host)
	if err != nil {
		log.Fatal("redis error: %v", err)
	}
	if pass != "" {
		_ , err = client.Do("AUTH", pass)
		if err != nil {
			log.Fatal("redis error: %v", err)
		}
	}
	return client
}
