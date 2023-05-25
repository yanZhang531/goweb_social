package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/pzlymformeet/social/pkg/models"
	"github.com/pzlymformeet/social/pkg/utils"
	"log"
	"net/http"
	"strconv"
	"time"
)

/*
1. 获取所有的动态数据 http://127.0.0.1:8080/posts 		GET
2. 发布一条动态      http://127.0.0.1:8080/posts 		POST 携带动态信息
3. 删除一条动态	    http://127.0.0.1:8080/posts/1	DELETE
*/

func DispatchPosts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getPosts(w, r)
	case "POST":
		createPost(w, r)
	case "DELETE":
		DeletePost(w, r)
	default:
		http.Error(w, "Method is wrong!", http.StatusMethodNotAllowed)
		return
	}
}

// 1. 获取所有的动态数据 http://127.0.0.1:8080/posts?userID=12345 		GET
func getPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.PostResult
	// 先根据查询参数获取userID，若为空则取所有人的动态数据
	err := r.ParseForm()
	if err != nil {
		log.Println("获取Form数据失败！")
		return
	}
	userID := r.Form.Get("userID")
	fmt.Println("userID:", userID)
	if userID == "" {
		posts, err = models.GetAllPosts()
		if err != nil {
			posts, err = models.GetUserAllPosts(userID)
			if err != nil {
				log.Println("err:", err)
				utils.Error(w, "获取动态数据失败！")
				return
			}
		}
	}

	if userID != "" {
		posts, err = models.GetUserAllPosts(userID)
		if err != nil {
			log.Println("err:", err)
			utils.Error(w, "该用户没有发表过文章！")
			return
		}
	}
	utils.Success(w, posts)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	user, isExist := utils.CheckToken(r)
	if !isExist {
		utils.Error(w, "创建动态失败！请登录后再创建！")
		return
	}

	post := models.Post{}
	body, err := utils.ParseBody(r)
	if err != nil {
		log.Println("createPost:err:", err)
		utils.Error(w, "创建动态失败！")
		return
	}

	err = json.Unmarshal(body, &post)
	if err != nil {
		log.Println("createPost:Unmarshal:err:", err)
		utils.Error(w, "创建动态失败！")
		return
	}

	post.UserId = user.ID
	createTime := time.Now().Format("2006-01-02 15:04")
	post.CreateAt = createTime
	p, err := models.CreatePost(post)
	if err != nil {
		log.Println("createPost:CreatePost:err:", err)
		utils.Error(w, "创建动态失败！")
		return
	}
	utils.Success(w, p)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	user, isExist := utils.CheckToken(r)
	if !isExist {
		utils.Error(w, "创建动态失败！请登录后再创建！")
		return
	}
	postId := utils.ParsePath(r)

	userId := strconv.Itoa(user.ID)
	fmt.Println(postId, userId)
	err := models.DelPost(postId, userId)
	if err != nil {
		log.Println("DeletePost:DelPost :err:", err)
		utils.Error(w, "删除动态失败！")
		return
	}
	utils.Success(w, "删除成功！")
}
