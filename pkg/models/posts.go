package models

import (
	"errors"
)

type Post struct {
	ID          int    `json:"id" form:"id"`
	Description string `json:"description" form:"description"`
	Img         string `json:"img" form:"img"`
	UserId      int    `json:"userId" form:"userId"`
	CreateAt    string `json:"createAt" form:"createAt"'"`
}

type PostResult struct {
	Post
	Name       string `json:"name" form:"name"`
	ProfilePic string `json:"profile_pic" form:"profile_pic"`
}

// 根据动态id获取动态信息
func GetPostById(postId string) (*Post, error) {
	sql := "select * from posts where id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	post := Post{}
	err = stmt.QueryRow(postId).Scan(&post.ID, &post.Description, &post.Img, &post.UserId, &post.CreateAt)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

// 获取指定用户的全部动态信息
func GetUserAllPosts(userId string) ([]PostResult, error) {
	sql := "select p.*,u.name,u.profilePic from posts as p join users as u on (p.userId = u.id) where p.userId = ? order by p.createAt desc"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	pResult := make([]PostResult, 0)
	rows, err := stmt.Query(userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		r := PostResult{}
		rows.Scan(&r.ID, &r.Description, &r.Img, &r.UserId, &r.CreateAt, &r.Name, &r.ProfilePic)
		pResult = append(pResult, r)
	}
	if len(pResult) == 0 {
		return nil, errors.New("can't find posts!")
	}
	return pResult, nil
}

// 获取所有用户的动态信息
func GetAllPosts() ([]PostResult, error) {
	sql := "select p.*,u.name,u.profilePic from posts as p join users as u on (p.userId = u.id) order by p.createAt desc"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	pResult := make([]PostResult, 0)
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		r := PostResult{}
		rows.Scan(&r.ID, &r.Description, &r.Img, &r.UserId, &r.CreateAt, &r.Name, &r.ProfilePic)
		pResult = append(pResult, r)
	}
	return pResult, nil
}

// 发布动态
func CreatePost(post Post) (*Post, error) {
	sql := "insert into posts(description,img,userId,createAt) values (?,?,?,?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(post.Description, post.Img, post.UserId, post.CreateAt)
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	post.ID = int(id)
	return &post, nil
}

// 删除动态,根据动态Id和用户Id删除
func DelPost(pid string, userId string) error {
	sql := "delete from posts where id=? and userId = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(pid, userId)
	if err != nil {
		return err
	}
	affected, _ := result.RowsAffected()
	if affected != 1 {
		return errors.New("can't find post!")
	}
	return nil
}
