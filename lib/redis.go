package lib

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	redis_host   string
	password     string
	port     	 int
	maxOpenConns = 2
	maxIdleConns = 2
)

var Cache *RedisConnPool

type RedisConnPool struct {
	redisPool *redis.Pool
}



func init() {
	runMode := beego.BConfig.RunMode
	redis_host = beego.AppConfig.String(runMode + "::redis_host")
	password = beego.AppConfig.String(runMode + "::redis_password")
	port, _ = beego.AppConfig.Int(runMode + "::redis_port")

	Cache = new(RedisConnPool)
	Cache.redisPool = myNewPool()
	if Cache.redisPool == nil {
		panic("init redis failed！")
	}
}

func myNewPool() *redis.Pool {
	return &redis.Pool{
		MaxActive:   maxOpenConns,
		MaxIdle:     maxIdleConns,
		IdleTimeout: 10 * time.Second,
		Dial: func() (redis.Conn, error) {
			link := fmt.Sprintf("%s:%d",redis_host,port)
			c, err := redis.Dial("tcp", link)
			if err != nil {
				return nil, err
			}
			_, err = c.Do("AUTH", password)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func RedisDo(command string, args ...interface{}) (interface{}, error) {
	connect := Cache.redisPool.Get()
	defer connect.Close()
	return connect.Do(command, args...)
}

func RedisSetString(key string, value string, expired int) (interface{}, error)  {
	connect := Cache.redisPool.Get()
	defer connect.Close()
	if expired != 0 {
		return connect.Do("SET", key, value, "EX", expired)
	}else{
		return connect.Do("SET", key, value)
	}
}
//
func RedisGetString(key string) (string, error)  {
	connect := Cache.redisPool.Get()
	defer connect.Close()

	return redis.String(connect.Do("GET", key))
}