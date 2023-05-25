package routers

import (
	"github.com/pzlymformeet/social/pkg/controllers"
	"github.com/pzlymformeet/social/pkg/utils"
	"net/http"
)

/**
	http://127.0.0.1:8090/likes?postId=19 	Get 获取某条动态下的所有收藏信息
	http://127.0.0.1:8090/likes 			Post 收藏某条动态
	http://127.0.0.1:8090/likes?postId=1 	Delete 删除收藏
**/

func registerLikeRouter(mux *http.ServeMux) {
	mux.HandleFunc("/likes/", utils.CORS(controllers.DispatchLikes))
}
