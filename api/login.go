package api

import (
	"errors"
	"fmt"
	"goblog/common"
	"goblog/dao"
	"goblog/models"
	"goblog/utils"
	"net/http"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	// 接受用户名和密码 返回对应的json数据
	param := common.GetRequestJsonParam(r)
	username := param["username"].(string)
	passwd := param["passwd"].(string)
	fmt.Println(username)
	fmt.Println(passwd)
	passwd = utils.Md5Crypt(passwd, "mszlu")
	fmt.Println(passwd)
	loginReq := new(models.LoginReq)
	loginReq.Name = username
	loginReq.Passwd = passwd
	user, dbError := dao.Login(loginReq)
	if dbError != nil {
		if dbError.IsNilError {
			dbError.Err = errors.New("账号密码不正确")
		}
		common.Error(w, dbError.Err)
		return
	}
	uid := user.Uid
	token, _ := utils.Award(&uid)
	loginRes := &models.LoginResp{Token: token, UserInfo: models.UserRes{user.Uid, user.UserName, user.Avatar}}

	common.Success(w, loginRes)
}
