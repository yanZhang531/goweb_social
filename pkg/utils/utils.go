package utils

import (
	"encoding/json"
	"fmt"
	"github.com/pzlymformeet/social/pkg/models"
	"net/http"
	"strings"
)

func ParseBody(r *http.Request) ([]byte, error) {
	length := r.ContentLength
	data := make([]byte, length)
	_, err := r.Body.Read(data)
	if err != nil && err.Error() != "EOF" {
		return nil, err
	}
	return data, nil
}

func User2Map(user *models.User) map[string]any {
	data := map[string]any{
		"id":         user.ID,
		"username":   user.Username,
		"Email":      user.Email,
		"Name":       user.Name,
		"CoverPic":   user.CoverPic,
		"ProfilePic": user.ProfilePic,
		"City":       user.City,
		"WebSite":    user.WebSite,
	}
	return data
}

func ParsePath(r *http.Request) string {
	path := r.URL.Path
	params := strings.Split(path, "/")

	return params[len(params)-1]
}

func CheckToken(r *http.Request) (*models.User, bool) {
	token := r.Header.Get("Accesstoken")
	fmt.Println("cookie:", token)
	kvs := strings.Split(token, "=")
	if len(kvs) <= 1 {
		return nil, false
	}
	username := kvs[1]
	user, err := models.GetUserByUsername(username)
	if err != nil {
		return nil, false
	}
	return &user, true
}

func Error(w http.ResponseWriter, errInfo string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(errInfo))
}

func Success(w http.ResponseWriter, data any) {
	sendData, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(sendData)
}

func CORS(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("access-control-allow-origin", "*")
		w.Header().Set("access-control-allow-headers", "Content-type,accesstoken,x-xsrf-token,authorization,token")
		w.Header().Set("access-control-allow-credential", "true")
		w.Header().Set("access-control-allow-methods", "post,get,delete,put,options")
		w.Header().Set("access-type", "application/json;charset=utf-8")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		f(w, r)
	}
}
