package models

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUser_UpdateUserInfo(t *testing.T) {
	users := []User{
		{
			Username:   "23423",
			Password:   "132213",
			Email:      "42424",
			Name:       "赵美延",
			CoverPic:   "",
			ProfilePic: "",
			City:       "",
			WebSite:    "",
		}, {
			Username:   "0",
			Password:   "144324",
			Email:      "14342",
			Name:       "宋雨琦",
			CoverPic:   "",
			ProfilePic: "",
			City:       "",
			WebSite:    "",
		}, {
			Username:   "0",
			Password:   "432423",
			Email:      "2342342",
			Name:       "",
			CoverPic:   "",
			ProfilePic: "",
			City:       "",
			WebSite:    "",
		},
	}
	bytes, _ := json.Marshal(users)
	fmt.Println(string(bytes))
}
