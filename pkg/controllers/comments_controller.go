package controllers

import (
	"encoding/json"
	"github.com/pzlymformeet/social/pkg/models"
	"github.com/pzlymformeet/social/pkg/utils"
	"log"
	"net/http"
)

func DispatchComments(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getComments(w, r)
	case "POST":
		addComment(w, r)
	default:
		utils.Error(w, "请求错误！")
		return
	}
}

// http://127.0.0.1:8010/comments?postId=20	GET：获取某条动态下所有的评论信息
func getComments(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("getComments:err:", err)
		utils.Error(w, "获取评论失败！")
		return
	}
	postId := r.Form.Get("postId")
	comments, err := models.GetPostComments(postId)
	if err != nil {
		log.Println("getComments:err:", err)
		utils.Error(w, "获取评论失败！")
		return
	}
	utils.Success(w, comments)
}

func addComment(w http.ResponseWriter, r *http.Request) {
	user, isExist := utils.CheckToken(r)
	if !isExist {
		log.Println("addComment:err:未登录！")
		utils.Error(w, "添加评论失败！请登录！")
		return
	}
	comment := models.Comment{}
	body, err := utils.ParseBody(r)
	if err != nil {
		log.Println("addComment:err:", err)
		utils.Error(w, "添加评论失败！")
		return
	}
	err = json.Unmarshal(body, &comment)
	if err != nil {
		log.Println("addComment:err:", err)
		utils.Error(w, "添加评论失败！")
		return
	}
	comment.UserId = user.ID
	c, err := models.AddComment(comment)
	if err != nil {
		log.Println("addComment:err:", err)
		utils.Error(w, "添加评论失败！")
		return
	}
	utils.Success(w, c)
}
