package main

import (
	"github.com/gin-gonic/gin"
	"GinWebService/model"
	"fmt"
	"GinWebService/controller"
)

func main(){
	//初始化路由
	router:=gin.Default()
	//请求分配
	model.InitRedis()
	err:=model.InitDb()

	if err != nil {
		fmt.Println(err)
		return
	}
	router.Static("/home","view")

	r1:=router.Group("/api/v1.0")
	{
		//路由规范
		r1.GET("/areas",controller.GetArea)
	}
	router.Run(":8082")
}