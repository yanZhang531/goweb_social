package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/pzlymformeet/social/pkg/models"
	"github.com/pzlymformeet/social/pkg/utils"
	"log"
	"net/http"
	"strconv"
)

func DispatchRelationship(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getRelationship(w, r)
	case "POST":
		addRelationship(w, r)
	case "PUT":
	case "DELETE":
		deleteRelationship(w, r)
	default:
		http.Error(w, "wrong Method!", http.StatusMethodNotAllowed)
		return
	}

}

// 查找关注的人
// http://127.0.0.1:8010/relationship?userId=123123
func getRelationship(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// r.URL.Query().Get("userId")
	id := r.FormValue("userId")
	followerId, _ := strconv.Atoi(id)

	relationships, err := models.GetFollowed(followerId)
	if err != nil {
		http.Error(w, "there no relationships", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	bytes, _ := json.Marshal(relationships)
	w.Write(bytes)

}

func addRelationship(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// 验证是否是登录状态，只有登录状态才能关注
	loginUser, isLogin := utils.CheckToken(r)
	if !isLogin {
		http.Error(w, "Not log in!please log in!", http.StatusBadRequest)
		return
	}
	// 从body中获取要关注的id
	body, err := utils.ParseBody(r)
	if err != nil {
		http.Error(w, "can't parse the body!", http.StatusBadRequest)
		return
	}
	relation := models.Relationship{}
	err = json.Unmarshal(body, &relation)
	if err != nil {
		http.Error(w, "can't unmarshal the body!", http.StatusBadRequest)
		return
	}
	relation.FollowerUserId = loginUser.ID
	// 在数据库中添加关系
	err = relation.CreateRelationship()
	if err != nil {
		log.Println(err)
		http.Error(w, "can't add the relationship!", http.StatusBadRequest)
		return
	}
	// 返回关注信息
	b, err := json.Marshal(relation)
	if err != nil {
		http.Error(w, "can't unmarshal the body!", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// http://127.0.0.1:8010/relationship?userId=123123
func deleteRelationship(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// 是否登录
	loginUser, isLogin := utils.CheckToken(r)
	if !isLogin {
		http.Error(w, "Not log in!please log in!", http.StatusBadRequest)
		return
	}

	// r.URL.Query().Get("userId")
	id := r.FormValue("userId") // 取关的用户号
	followedId, _ := strconv.Atoi(id)

	relation := models.Relationship{
		FollowedUserId: followedId,
		FollowerUserId: loginUser.ID,
	}
	err := relation.DelRelationship()
	if err != nil {
		http.Error(w, "unfollowed failed!", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "unfollowed success!")
}
