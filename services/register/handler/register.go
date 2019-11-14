package handler

import (
	"context"
	register "GinWebService/services/register/proto/register"
	"GinWebService/services/register/utils"
	"errors"
	"GinWebService/services/register/model"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"math/rand"
	"fmt"
	"time"
	"crypto/md5"
	"encoding/hex"
)

type Register struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Register) SmsCode(ctx context.Context, req *register.Request, rsp *register.Response) error {
	//验证图片验证码是否输入正确
	rnd, err := model.GetImgCode(req.Uuid)
	if err != nil {
		rsp.Errno = utils.RECODE_NODATA
		rsp.Errmsg = utils.RecodeText(utils.RECODE_NODATA)
		return err
	}
	if req.Text != rnd {
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		return errors.New("验证码输入错误")
	}
	//如果成功,发送短信,存储短信验证码  阿里云短信接口
	client, err := sdk.NewClientWithAccessKey("default", "LTAI4FexwrAFbn4ua4DHAyXh", "AltI2inQ1I5TqAEwAfrJNgP54VnVOx")
	if err != nil {
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		return err
	}
	//获取6位数随机码
	myRnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06d", myRnd.Int31n(1000000))

	//初始化请求对象
	request := requests.NewCommonRequest()
	request.Method = "POST"                                         //设置请求方法
	request.Scheme = "https"                                        // https | http   //设置请求协议
	request.Domain = "dysmsapi.aliyuncs.com"                        //域名
	request.Version = "2017-05-25"                                  //版本号
	request.ApiName = "SendSms"                                     //api名称
	request.QueryParams["PhoneNumbers"] = req.Mobile                //需要发送的电话号码
	request.QueryParams["SignName"] = "北京5期区块链"                     //签名名称   需要申请
	request.QueryParams["TemplateCode"] = "SMS_176375357"           //模板号   需要申请
	request.QueryParams["TemplateParam"] = `{"code":` + vcode + `}` //发送短信验证码

	response, err := client.ProcessCommonRequest(request) //发送短信
	//如果不成功
	if !response.IsSuccess() {
		rsp.Errno = utils.RECODE_SMSERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_SMSERR)
		return errors.New("发送短信失败")
	}
	//存储短信验证码,存入redis中
	err = model.SaveSmsCode(req.Mobile, vcode)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	return nil
}

//注册
func (e *Register) Register(ctx context.Context, req *register.RegRequset, rsp *register.RegResponse) error {
	//把数据存储到mysql中   校验短信 验证码是否正确
	smsCode, err := model.GetSmsCode(req.Mobile)
	if err != nil {
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		return err
	}
	if smsCode != req.SmsCode {
		rsp.Errno = utils.RECODE_SMSERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_SMSERR)
		return errors.New("验证码错误")
	}
	//存储用户数据
	//给密码加密
	m5 := md5.New()
	m5.Write([]byte(req.Password))
	pwd_hash := hex.EncodeToString(m5.Sum(nil))
	err = model.SaveUser(req.Mobile, pwd_hash)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	return nil
}

//登录
func (e *Register) Login(ctx context.Context, req *register.RegRequset, rsp *register.RegResponse) error {
	//加密密码
	m5 := md5.New()
	m5.Write([]byte(req.Password))
	pwd_hash := hex.EncodeToString(m5.Sum(nil))

	user, err := model.CheckUser(req.Mobile, pwd_hash)
	if err != nil {
		rsp.Errno = utils.RECODE_LOGINERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_LOGINERR)
		return err
	}
	//查询成功,将用户名存入session
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	rsp.Name=user.Name
	return nil
}
