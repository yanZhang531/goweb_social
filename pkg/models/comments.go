package models

import (
	"errors"
	"time"
)

type Comment struct {
	ID          int    `json:"id" form:"id"`
	Description string `json:"description" form:"description"`
	CreateAt    string `json:"createAt" form:"createAt"'"`
	UserId      int    `json:"userId" form:"userId"`
	PostId      int    `json:"postId" form:"postId"`
}

type CommentResult struct {
	Comment
	ProfilePic string `json:"profilePic" form:"profilePic"`
	Name       string `json:"name" form:"name"`
}

// 获取某条动态的所有评论
func GetPostComments(postId string) ([]CommentResult, error) {
	sql := "select c.*,u.name,u.profilePic from comments as c join users as u on (c.userid = u.id) where postid = ? order by c.createAt desc"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	result := make([]CommentResult, 0)
	rows, err := stmt.Query(postId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		r := CommentResult{}
		err = rows.Scan(&r.ID, &r.Description, &r.CreateAt, &r.UserId, &r.PostId, &r.Name, &r.ProfilePic)
		if err != nil {
			return nil, err
		}
		result = append(result, r)
	}
	return result, nil
}

// 给某条动态添加评论
func AddComment(comment Comment) (*Comment, error) {
	sql := "insert into comments(description,userid,postid,createAt) values (?,?,?,?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	ctime := time.Now().Format("2006-01-02 15:04")
	result, err := stmt.Exec(comment.Description, comment.UserId, comment.PostId, ctime)
	if err != nil {
		return nil, err
	}
	affectedRow, err := result.RowsAffected()
	if affectedRow != 1 || err != nil {
		return nil, errors.New("insert comment failed")
	}
	id, _ := result.LastInsertId()
	comment.ID = int(id)
	return &comment, nil
}
