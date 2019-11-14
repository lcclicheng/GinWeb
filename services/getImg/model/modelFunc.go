package model

import "github.com/garyburd/redigo/redis"

var RedisPool redis.Pool
//redis连接池
func InitRedis(){
	RedisPool=redis.Pool{
		MaxIdle:100,
		MaxActive:100,
		IdleTimeout:60*5,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp","127.0.0.1:6379")
		},
	}
}
//存储验证码
func SaveImgRnd(uuid,rnd string)error{
	conn:=RedisPool.Get()
	_,err:=conn.Do("setex",uuid,60*5,rnd)
	return err
}