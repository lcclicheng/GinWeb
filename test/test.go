package main

import (
	"github.com/afocus/captcha"
	"image/color"
	"net/http"
	"image/png"
)

func main(){
	cap:= captcha.New()
	//设置字符集
	if err:=cap.SetFont("comic.ttf"); err!=nil{
		panic(err.Error())
	}
	//设置验证码图片大小
	cap.SetSize(128,64)
	//设置混淆程度
	cap.SetDisturbance(captcha.MEDIUM)
	//设置字体颜色
	cap.SetFrontColor(color.RGBA{255,255,255,255})
	//设置背景色
	cap.SetBkgColor(color.RGBA{255,0,0,255},color.RGBA{0,0,255,255},color.RGBA{0,153,0,255})
	//创建验证码图片
	http.HandleFunc("/r",func(w http.ResponseWriter,r *http.Request){
		img,str:=cap.Create(2,captcha.NUM)
		png.Encode(w,img)
		println(str)
	})
	http.HandleFunc("/c",func(w http.ResponseWriter,r *http.Request){
		str:=r.URL.RawQuery
		img:=cap.CreateCustom(str)
		png.Encode(w,img)
	})
	http.ListenAndServe(":8810",nil)
}