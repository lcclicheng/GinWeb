package handler

import (
	"context"
	user "GinWebService/services/user/proto/user"
	"GinWebService/services/user/model"

	"GinWebService/services/user/utils"
)

type User struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *User) MicroGetUser(ctx context.Context, req *user.Request, rsp *user.Response) error {
	//根据用户名获取用户信息  ,从mysql数据库中获取
	myUser,err:=model.GetUserInfo(req.Name)
	if err != nil {
		rsp.Errno=utils.RECODE_USERERR
		rsp.Errmsg=utils.RecodeText(utils.RECODE_USERONERR)
		return err
	}
	rsp.Errno=utils.RECODE_OK
	rsp.Errmsg=utils.RecodeText(utils.RECODE_OK)

	//获取一个结构体
	var userInfo user.UserInfo
	userInfo.UserId=int32(myUser.ID)
	userInfo.Name=myUser.Name
	userInfo.Mobile=myUser.Mobile
	userInfo.RealName=myUser.Real_name
	userInfo.IdCard=myUser.Id_card
	userInfo.AvatarUrl=myUser.Avatar_url

	rsp.Data=&userInfo
	return nil
}

func (e *User)UpdateUserName(ctx context.Context, req *user.UpdateReq, resp*user.UpdateResp) error{
	err:=model.UpdateUserName(req.OldName,req.NewName)
	if err != nil {
		resp.Errno=utils.RECODE_DATAERR
		resp.Errmsg=utils.RecodeText(utils.RECODE_DATAERR)
		return nil
	}
	resp.Errno=utils.RECODE_OK
	resp.Errmsg=utils.RecodeText(utils.RECODE_OK)
	var nameData user.NameData
	nameData.Name=req.NewName
	resp.Data=&nameData
	return nil

}