package db

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

var RedisPool *redis.Pool

func InitRedis() {
	RedisPool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", ":6379")
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
}

func GetRedisString(key string) (string, error) {
	redisClient := RedisPool.Get()
	defer redisClient.Close()
	res, err := redis.String(redisClient.Do("Get", key))
	if err != nil {
		log.Println("redis err :", err)
		return "", err
	}
	return res, nil
}

func SetRedisString(params ...interface{}) (string, error) {
	redisClient := RedisPool.Get()
	defer redisClient.Close()
	res, err := redis.String(redisClient.Do("Set", params[0], params[1]))
	if err != nil {
		log.Println("redis err :", err)
		return "", err
	}
	return res, nil
}

func ExitsKey(key string) (int, error) {
	redisClient := RedisPool.Get()
	defer redisClient.Close()
	res, err := redis.Int(redisClient.Do("EXISTS", key))
	if err != nil {
		log.Println("redis err :", err)
		return 0, err
	}
	return res, nil
}

func Incr(key string) {
	redisClient := RedisPool.Get()
	defer redisClient.Close()
	redisClient.Do("INCR", key)
}
