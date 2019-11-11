package controller

import (
	"github.com/gin-gonic/gin"
    getArea "GinWebService/proto/getArea/proto/getArea"
	"github.com/micro/go-micro/client"
	"context"
	"fmt"
	"net/http"
)
//获取所有地区信息
func GetArea (ctx  *gin.Context){
	/*resp:=make(map[string]interface{})
	areas,err:=model.GetArea()
	if err != nil {
		fmt.Println("获取所有地域信息失败")
		resp["errno"]=utils.RECODE_DBERR
		resp["errmsg"]=utils.RecodeText(utils.RECODE_DBERR)
		ctx.JSON(http.StatusOK,resp)
		return
	}
	//把数据返回给前端
	resp["errno"]=utils.RECODE_OK
	resp["errmsg"]=utils.RecodeText(utils.RECODE_OK)
	resp["data"]=areas
	ctx.JSON(http.StatusOK,resp)*/
	//调用远程服务,获取所有地域信息
	microClient:=getArea.NewGetAreaService("go.micro.srv.getArea",client.DefaultClient)
	resp,err:=microClient.MicroGetArea(context.TODO(),&getArea.Request{})
	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(http.StatusOK,resp)
}