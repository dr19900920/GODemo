package controllers

import (
	"common"
	"models"
	"net/http"
)

/**
* 用户请求
 */
func Login(w http.ResponseWriter, r *http.Request) {
	//解析表单
	r.ParseForm()
	//登录邮箱
	email := r.FormValue("email")
	//登录密码
	password := r.FormValue("password")

	if email == "" {
		w.Write(common.NewResponseSimple(400, "请输入邮箱").Encode())
	} else if password == "" {
		w.Write(common.NewResponseSimple(400, "请输入密码").Encode())
	} else {
		num, err := models.CheckAccount(email)
		if err != nil {
			w.Write(common.NewResponseSimple(400, err.Error()).Encode())
			return
		}

		if num > 0 {
			//验证密码是否正确
			//密码正确的情况, 获取用户信息
			user, err := models.GetUserInfo(email)
			if err != nil {
				w.Write(common.NewResponseSimple(400, err.Error()).Encode())
				return
			}
			w.Write(common.NewResponseData(user, "用户登录成功").Encode())
		} else {
			w.Write(common.NewResponseSimple(400, "用户不存在").Encode())
		}
	}
}
