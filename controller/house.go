package controller

import (
	"github.com/gin-gonic/gin"
    getArea "GinWebService/proto/getArea/proto/getArea"
    getImg "GinWebService/proto/getImg"
	"context"
	"fmt"
	"net/http"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	"GinWebService/utils"
	"image/png"
	"github.com/afocus/captcha"
	"encoding/json"
	register "GinWebService/proto/register"
	"github.com/gin-contrib/sessions"
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
	consulRegistry:=consul.NewRegistry()
	microService:=micro.NewService(
		micro.Registry(consulRegistry),
	)
	//调用远程服务,获取所有地域信息
	microClient:=getArea.NewGetAreaService("go.micro.srv.getArea",microService.Client())
	resp,err:=microClient.MicroGetArea(context.TODO(),&getArea.Request{})
	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(http.StatusOK,resp)
}

//写一个假的session
func GetSession(ctx *gin.Context){
	resp:=make(map[string]interface{})

	resp["errno"]=utils.RECODE_LOGINERR
	resp["errmsg"]=utils.RecodeText(utils.RECODE_LOGINERR)
	ctx.JSON(http.StatusOK,resp)
}
//获取验证码图片方法
func GetImageCd(ctx *gin.Context){
	//获取数据
	uuid:=ctx.Param("uuid")
	//校验数据
	if uuid==""{
		fmt.Println("获取数据错误")
		return
	}
	//处理数据
	//调用远程服务
	consulReg:=consul.NewRegistry()
	microService:=micro.NewService(
		micro.Registry(consulReg),
	)
	microClient:=getImg.NewGetImgService("go.micro.srv.getImg",microService.Client())
	resp,err:=microClient.MicroGetImg(context.TODO(),&getImg.Request{Uuid:uuid})
	//获取数据
	if err != nil {
		fmt.Println("---->",err)
		fmt.Println("获取远端数据失败")
		ctx.JSON(http.StatusOK,resp)
		return
	}
	//返回json数据
	var img captcha.Image
	json.Unmarshal(resp.Data,&img)
	png.Encode(ctx.Writer,img)

}

//获取短信验证码
func GetSmsCd(ctx *gin.Context){
	//获取数据
	mobile:=ctx.Param("mobile")
	//输入的图片验证码
	text:=ctx.Query("text")
	//获取验证码图片的uuid
	uuid:=ctx.Query("id")
	if mobile==""||text==""||uuid==""{
		fmt.Println("传入数据不完整")
		return
	}
    //初始化客户端
    microClient:=register.NewRegisterService("go.micro.srv.register",utils.GetMicroClient())
    //调用远程客户端
    resp,err:=microClient.SmsCode(context.TODO(),&register.Request{
    	Uuid:uuid,
    	Text:text,
    	Mobile:mobile,
    })
	if err != nil {
		fmt.Println("===>",err)
		fmt.Println("调用远程服务错误",err)
	}
	ctx.JSON(http.StatusOK,resp)
}

//注册方法

type RegStu struct {
	Mobile string `json:"mobile"`
	PassWord  string`json:"password"`
	SmsCode  string`json:"sms_code"`
}
//注册业务
func PostRet(ctx*gin.Context){
	//获取数据
	//mobile:=ctx.PostForm("mobile")
	//pwd:=ctx.PostForm("password")
	//smsCode:=ctx.PostForm("sms_code")
	var reg RegStu
	err:=ctx.Bind(&reg)
	//校验数据
	if err != nil {
		fmt.Println("获取前端传递数据失败")
		return
	}
	//处理数据
	microClient:=register.NewRegisterService("go.micro.srv.register",utils.GetMicroClient())
	resp,err:=microClient.Register(context.TODO(),&register.RegRequset{
		Mobile:reg.Mobile,
		Password:reg.PassWord,
		SmsCode:reg.SmsCode,
	})
	if err != nil {
		fmt.Println("调用远程服务失败",err)
	}
	//返回数据
	ctx.JSON(http.StatusOK,resp)
}
type LogStu struct {
	Mobile string `json:"mobile"`
	PassWord  string`json:"password"`
}
//登录业务
func PostLogin(ctx *gin.Context){
	//获取数据
	var log LogStu
	err:=ctx.Bind(&log)
	//校验数据
	if err != nil {
		fmt.Println("获取数据失败")
		return
	}
	//处理数据
	microClient:=register.NewRegisterService("go.micro.srv.register",utils.GetMicroClient())
	resp,err:=microClient.Login(context.TODO(),&register.RegRequset{Mobile:log.Mobile,Password:log.PassWord})
	defer ctx.JSON(http.StatusOK,resp)
	if err != nil {
		fmt.Println("调用login错误",err)
		return
	}
	//返回数据,存储session,并返回数据给web端
	session:=sessions.Default(ctx)
	session.Set("userName",resp.Name)
	session.Save()
}