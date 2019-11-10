package main

import "github.com/gin-gonic/gin"

func main(){
	//初始化路由
	router:=gin.Default()
	//请求分配
	r1:=router.Group("/abc")
	{
		r1.GET("/efg", func(ctx *gin.Context) {
			ctx.Writer.WriteString("hello")
		})
	}
	router.Run(":8081")
}