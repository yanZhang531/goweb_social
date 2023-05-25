package controllers

import (
	"encoding/json"
	"github.com/pzlymformeet/social/pkg/models"
	"github.com/pzlymformeet/social/pkg/utils"
	"log"
	"net/http"
	"strconv"
)

func DispatchLikes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getLikes(w, r)
	case "POST":
		addLikes(w, r)
	case "DELETE":
		deleteLikes(w, r)
	default:
		utils.Error(w, "请求错误！")
		return
	}
}

// http://localhost:8010/likes/?postId=7
func getLikes(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("getLikes:ParseForm:err:", err)
		utils.Error(w, "获取收藏失败！")
		return
	}
	postId := r.Form.Get("postId")
	likes, err := models.GetLikesByPostId(postId)
	if err != nil {
		log.Println("getLikes:GetLikesByPostId:err:", err)
		utils.Error(w, "获取收藏失败！")
		return
	}
	utils.Success(w, likes)
}

// http://localhost:8010/likes/
func addLikes(w http.ResponseWriter, r *http.Request) {
	user, isExist := utils.CheckToken(r)
	if !isExist {
		log.Println("addLikes:CheckToken!未登录！")
		utils.Error(w, "请登录！")
	}
	like := models.Like{}
	body, err := utils.ParseBody(r)
	if err != nil {
		log.Println("addLikes:ParseBody!err:", err)
		utils.Error(w, "添加收藏失败！")
	}
	like.UserId = user.ID
	err = json.Unmarshal(body, &like)
	if err != nil {
		log.Println("addLikes:Unmarshal！")
		utils.Error(w, "添加收藏失败！")
	}

	l, err := models.CreateLike(like)
	if err != nil {
		log.Println("addLikes:Unmarshal！")
		utils.Error(w, "添加收藏失败！")
		return
	}
	utils.Success(w, l)

}

// http://localhost:8010/likes/?postId=7
func deleteLikes(w http.ResponseWriter, r *http.Request) {
	user, isExist := utils.CheckToken(r)
	if !isExist {
		log.Println("addLikes:CheckToken!未登录！")
		utils.Error(w, "请登录！")
	}
	postIdStr := r.FormValue("postId")
	postId, _ := strconv.Atoi(postIdStr)
	like := models.Like{
		UserId: user.ID,
		PostId: postId,
	}
	err := models.DelLikes(like)
	if err != nil {
		log.Println("addLikes:Unmarshal！")
		utils.Error(w, "删除收藏失败！")
		return
	}
	utils.Success(w, "删除收藏成功！")
}
