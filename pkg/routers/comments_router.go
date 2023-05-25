package routers

import (
	"github.com/pzlymformeet/social/pkg/controllers"
	"github.com/pzlymformeet/social/pkg/utils"
	"net/http"
)

/*
*

	http://127.0.0.1:8010/comments?postId=20	GET：获取某条动态下所有的评论信息
	http://127.0.0.1:8010/comments	POST:给某条动态评论

*
*/
func RegisterCommentRouter(mux *http.ServeMux) {
	mux.HandleFunc("/comments/", utils.CORS(controllers.DispatchComments))
}
