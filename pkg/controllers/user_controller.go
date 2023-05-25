package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/pzlymformeet/social/pkg/models"
	"github.com/pzlymformeet/social/pkg/utils"
	"log"
	"net/http"
)

// http:127.0.0.1/fetch/1
func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not GET!", http.StatusMethodNotAllowed)
		return
	}

	// 从request中拿url
	username := utils.ParsePath(r)
	if username == "" {
		http.Error(w, "username is not valid!", http.StatusBadRequest)
		return
	}

	// 根据username查数据
	user, err := models.GetUserByUsername(username)
	if err != nil {
		log.Println(err)
		http.Error(w, "cannot get the user!", http.StatusBadRequest)
		return
	}

	// 返回用户信息
	u, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(u)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not POST!", http.StatusMethodNotAllowed)
		return
	}

	body, err := utils.ParseBody(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "can't parse the body!", http.StatusBadRequest)
		return
	}

	// 获得要更新的信息
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println(err)
		http.Error(w, "can't parse the body!", http.StatusBadRequest)
		return
	}
	fmt.Println("user:", user)

	// 判断用户是否已经登录
	cookieUser, isExist := utils.CheckToken(r)
	if !isExist {
		http.Error(w, "user is not login!", http.StatusBadRequest)
		return
	}
	if cookieUser.Username != user.Username {
		// 已登录的用户与要修改信息的用户不匹配
		http.Error(w, "update user is not login user!", http.StatusBadRequest)
		return
	}
	err = user.UpdateUserInfo()
	if err != nil {
		log.Println(err)
		http.Error(w, "update user info failed!", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	tmp, _ := json.Marshal(cookieUser)
	w.Write(tmp)
}
