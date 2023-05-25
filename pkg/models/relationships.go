package models

import (
	"errors"
	"log"
)

type Relationship struct {
	ID             int `json:"id" form:"id"`                             // 关系ID
	FollowedUserId int `json:"followed_user_id" form:"followed_user_id"` // 被关注者ID
	FollowerUserId int `json:"follower_user_id" form:"follower_user_id"` // 关注者ID
}

// 创建关系，xx关注了xx
func (r *Relationship) CreateRelationship() error {
	sql := "insert into relationships(FollowedUserId,FollowerUserId) values (?,?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(r.FollowedUserId, r.FollowerUserId)
	if err != nil {
		return err
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	r.ID = int(insertId)
	return nil
}

// 获取被关注的所有人，followerUserId关注了xx
func GetFollowed(followerUserId int) ([]Relationship, error) {
	sql := "select id,followedUserId,followerUserId from relationships where followerUserId = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(followerUserId)
	if err != nil {
		return nil, err
	}
	rs := make([]Relationship, 0)
	for rows.Next() {
		r := Relationship{}
		rows.Scan(&r.ID, &r.FollowedUserId, &r.FollowerUserId)
		rs = append(rs, r)
	}
	return rs, nil
}

// 获取关系，followedUserId的粉丝有哪些
func GetFollowers(followedUserId int) ([]Relationship, error) {
	sql := "select id,followedUserId,followerUserId from relationships where followedUserId = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(followedUserId)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rs := make([]Relationship, 0)
	for rows.Next() {
		r := Relationship{}
		rows.Scan(&r.ID, &r.FollowedUserId, &r.FollowerUserId)
		rs = append(rs, r)
	}
	return rs, nil
}

// DelRelationship 删除关系，取消关注
func (r *Relationship) DelRelationship() error {
	sql := "delete from relationships where followedUserId = ? and followerUserId = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(r.FollowedUserId, r.FollowerUserId)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil || affected != 1 {
		log.Println(err)
		return errors.New("can't delete!")
	}
	return nil
}
