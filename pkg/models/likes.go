package models

import "errors"

type Like struct {
	ID     int `json:"id" form:"id"`
	UserId int `json:"userId" form:"userId"`
	PostId int `json:"postId" form:"postId"`
}

// 用户收藏的所有帖子
func GetLikesByUserID() {

}

// 帖子的所有收藏人
func GetLikesByPostId(postId string) (likes []Like, err error) {
	sql := "select * from likes where postId = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(postId)
	if err != nil {
		return
	}
	for rows.Next() {
		like := Like{}
		err = rows.Scan(&like.ID, &like.UserId, &like.PostId)
		if err != nil {
			return
		}
		likes = append(likes, like)
	}
	return
}

func CreateLike(like Like) (*Like, error) {
	sql := "insert into likes(postId,userId) values (?,?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(like.PostId, like.UserId)
	if err != nil {
		return nil, err
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	like.ID = int(insertId)
	return &like, nil
}

func DelLikes(like Like) error {
	sql := "delete from likes where postId = ? and userId = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(like.PostId, like.UserId)
	if err != nil {
		return err
	}
	affectedRow, _ := result.RowsAffected()
	if affectedRow != 1 {
		return errors.New("can't delete like!")
	}
	return nil
}
