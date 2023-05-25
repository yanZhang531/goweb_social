package models

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"github.com/pzlymformeet/social/pkg/config"
	"log"
	"strings"
)

type User struct {
	ID         int    `json:"id" form:"id"`                   // 用户编号，非空
	Username   string `json:"username" form:"username"`       // 用户名，非空
	Password   string `json:"password" form:"password"`       // 密码，非空
	Email      string `json:"email" form:"email"`             // 邮箱，非空
	Name       string `json:"name" form:"form"`               // 昵称，非空
	CoverPic   string `json:"cover_pic" form:"cover_pic"`     //背景图
	ProfilePic string `json:"profile_pic" form:"profile_pic"` // 头像
	City       string `json:"city" form:"city"`               //城市
	WebSite    string `json:"web_site" form:"web_site"`       //个人网站
}

var db = config.GetDb()

// 获取所有用户
func GetAllUsers() (users []User, err error) {
	sql := "select id,username,email,name,coverpic,profilepic,city,website from users;"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()
	users, err = rows2Users(stmt)
	if err != nil {
		return
	}
	return
}

// 通过Id获取指定用户
func GetUserById(id int) (u User, err error) {
	// u的id不为空
	if id == 0 {
		return User{}, errors.New("cannot get the user! wrong id!")
	}
	sql := "select username,password,email,name,coverpic,profilepic,city,website from users where id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(sql, u.ID).Scan(&u.Username, &u.Password, &u.Email, &u.Name, &u.CoverPic, &u.ProfilePic, &u.City, &u.WebSite)
	if err != nil {
		return
	}
	return
}

// GetUserByUsername 根据username获取指定用户
func GetUserByUsername(username string) (user User, err error) {
	// u的id不为空
	if username == "" {
		return user, errors.New("cannot get the user! username is null!")
	}

	sql := "select id,password,username,email,name,coverpic,profilepic,city,website from users where username = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&user.ID,
		&user.Password, &user.Username, &user.Email, &user.Name, &user.CoverPic,
		&user.ProfilePic, &user.City, &user.WebSite)

	if err != nil {
		return
	}
	return
}

// 创建一个用户
func (u *User) CreateAUser() error {
	if u == nil {
		return errors.New("the user is nil!")
	}
	sql := "insert into users(username,password,email,name,coverpic,profilepic,city,website ) values(?,?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(u.Username, u.Password, u.Email, u.Name, u.CoverPic, u.ProfilePic, u.City, u.WebSite)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	u.ID = int(id)
	return nil
}

// 更新用户信息
func (u *User) UpdateUserInfo() error {
	if u.ID == 0 {
		return errors.New("ID is 0!")
	}
	if u.Username == "" {
		return errors.New("Username is nil!")
	}

	queryParams, sqlStrings := FindNotNull(u)
	sqlStr := strings.Join(sqlStrings, ",")
	sqlStr = strings.Replace(sqlStr, ",", " ", 1)
	finalStr := sqlStr + " where id = ?"
	log.Println("要更新的语句是：", finalStr)
	log.Println("替换的内容是：", queryParams)

	stmt, err := db.Prepare(finalStr)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(queryParams...)
	if err != nil {
		return errors.New("update wrong!")
	}

	return nil
}

func rows2Users(stmt *sql.Stmt) (users []User, err error) {
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := User{}
		rows.Scan(&user.ID, &user.Username, &user.Email, &user.Name, &user.CoverPic, &user.ProfilePic, &user.City, &user.WebSite)
		users = append(users, user)
	}
	return
}

func FindNotNull(u *User) ([]any, []string) {
	sqlStorage := make([]string, 0)
	params := make([]any, 0)

	tmpsql := "update users set"
	sqlStorage = append(sqlStorage, tmpsql)
	if u.Name != "" {
		sqlStorage = append(sqlStorage, "name = ?")
		params = append(params, u.Name)
	}
	if u.City != "" {
		sqlStorage = append(sqlStorage, "city = ?")
		params = append(params, u.City)
	}
	if u.CoverPic != "" {
		sqlStorage = append(sqlStorage, "coverpic = ?")
		params = append(params, u.CoverPic)
	}
	if u.WebSite != "" {
		sqlStorage = append(sqlStorage, "website = ?")
		params = append(params, u.WebSite)
	}
	if u.Email != "" {
		sqlStorage = append(sqlStorage, "email = ?")
		params = append(params, u.Email)
	}
	if u.Password != "" {
		sqlStorage = append(sqlStorage, "password = ?")
		pwd := EncryptPassword(u.Password)
		params = append(params, pwd)
	}
	if u.ProfilePic != "" {
		sqlStorage = append(sqlStorage, "profilepic = ?")
		params = append(params, u.ProfilePic)
	}
	params = append(params, u.ID)
	return params, sqlStorage
}

// 对密码做md5加密
func EncryptPassword(pwd string) string {
	md := md5.New()
	md.Write([]byte(pwd))
	return hex.EncodeToString(md.Sum(nil))
}

func VerifyPassword(pwd string, enPwd string) bool {
	return enPwd == EncryptPassword(pwd)
}
