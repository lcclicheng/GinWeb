package main

import (
	"github.com/gin-gonic/gin"
	"GinWebService/model"
	"fmt"
	"GinWebService/controller"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-contrib/sessions"
)

func main(){
	//初始化路由
	router:=gin.Default()
	//请求分配,数据库处理
	model.InitRedis()
	err:=model.InitDb()

	if err != nil {
		fmt.Println(err)
		return
	}
	//初始化redis容器,存储session数据
	store,_:=redis.NewStore(20,"tcp","127.0.0.1:6379","",[]byte("session"))
	store.Options(
		sessions.Options{
			MaxAge:0,
		},
	)
	router.Static("/home","view")
	/*router.Use()
	store,err:=redis.NewStore(20,"tcp","127.0.0.1:6379","",[]byte("session"))
	if err != nil {
		fmt.Println("初始化session容器错误")
		return
	}
	//路由使用中间件
	router.Use(sessions.Sessions("mysession",store))
	//使用路由就可以使用session中间件
	router.GET("/session", func(ctx *gin.Context) {
		//初始化session对象
		se:=sessions.Default(ctx)
		//设置session的时候,除了set之外,必须调用save
		se.Set("test","bj5q")
		se.Save()
		ctx.Writer.WriteString("设置session")
	})
	//获取session
	router.GET("/getSession", func(ctx *gin.Context) {
		//初始化session对象
		se:=sessions.Default(ctx)
		//获取session
		result:=se.Get("test")
		fmt.Println("得到的session为",result.(string))
		ctx.Writer.WriteString("获取session")
	})*/
	r1:=router.Group("/api/v1.0")
	{
		//路由规范
		r1.GET("/areas",controller.GetArea)
		r1.GET("/session",controller.GetSession)
		r1.GET("/imagecode/:uuid",controller.GetImageCd)
		r1.GET("/smscode/:mobile",controller.GetSmsCd)
		r1.POST("/users",controller.PostRet)
		r1.Use(sessions.Sessions("mysession",store))
		//登录业务
		r1.POST("/session",controller.PostLogin)
	}
	router.Run(":8080")
}