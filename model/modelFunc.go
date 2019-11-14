package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//获取所有地域信息

func GetArea()([]Area,error){
	//连接数据库
	var areas []Area

	//从redis中获取数据
	//redis.Dial("tcp","192.168.150.11:6379")
	conn:=GlobalRedis.Get()
	defer conn.Close()
	areaByte,_:=redis.Bytes(conn.Do("get","areaData"))
	if len(areaByte)==0{
		//从mysql拿数据
		if err :=GlobalDB.Find(&areas).Error;err!=nil {
			fmt.Println("err--find",err)
			return areas,err
		}
		//序列化数据存入redis
		//把数据存入redis
		areajson,err:=json.Marshal(areas)
		if err != nil {
			return nil,err
		}
		_,err=conn.Do("set","areaData",areajson)
		fmt.Println("conn.Do===>",err)
		fmt.Println("从mysql中获取数据")
	}else {
		json.Unmarshal(areaByte,&areas)
		fmt.Println("从redis中获取数据")
	}
	return areas,nil
}