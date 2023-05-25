package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/pzlymformeet/social/pkg/models"
	"github.com/pzlymformeet/social/pkg/utils"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// 判断是否是post方法
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 解析r中的username和password字段
	username := r.PostFormValue("username")
	pwd := r.PostFormValue("password")

	// 验证username是否在数据库中
	user, err := models.GetUserByUsername(username)
	if err != nil {
		// 未找到，查找用户出错
		http.Error(w, "Cannot find the User!", http.StatusBadRequest)
		return
	}

	// 验证password是否正确
	if !models.VerifyPassword(pwd, user.Password) {
		//密码不正确
		http.Error(w, "Password is not right!", http.StatusBadRequest)
		return
	}

	// 设置登录cookie
	cookie := http.Cookie{
		Name:  "username",
		Value: user.Username,
	}
	w.Header().Set("AccessToken", cookie.String())

	// 通过验证，返回用户信息
	data := utils.User2Map(&user)
	bytes, err := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "something was wrong!\n")
	}

}

func Register(w http.ResponseWriter, r *http.Request) {
	// 判断是否是post方法
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 解析r中的username和password字段
	registerUser := models.User{}
	storage, err := utils.ParseBody(r)
	if err != nil {
		log.Println("can't parse the body,err :", err)
		return
	}
	err = json.Unmarshal(storage, &registerUser)
	if err != nil {
		log.Println("can't unmarshal the data,err :", err)
		return
	}

	// 验证username是否在数据库中
	dataUser, err := models.GetUserByUsername(registerUser.Username)
	if err == nil && dataUser.Username == registerUser.Username {
		// 数据库中已经出现该用户账号，则不能注册该用户账号
		http.Error(w, "User is exists!Please change the username!", http.StatusBadRequest)
		return
	}

	// 完善用户信息
	registerUser.Password = models.EncryptPassword(registerUser.Password)
	registerUser.CoverPic = "../pkg/sources/default_cover_pic"
	registerUser.ProfilePic = "../pkg/sources/default_profile_pic"

	// 在数据库中创建用户
	err = registerUser.CreateAUser()
	if err != nil {
		http.Error(w, "Create a User failed!", http.StatusBadRequest)
		return
	}

	// 通过验证，返回用户信息
	data := utils.User2Map(&registerUser)
	bytes, err := json.Marshal(data)
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("accessToken", "")
	fmt.Fprintf(w, "成功退出！\n")
}
