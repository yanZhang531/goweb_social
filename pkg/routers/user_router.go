package routers

import (
	"github.com/pzlymformeet/social/pkg/controllers"
	"github.com/pzlymformeet/social/pkg/utils"
	"net/http"
)

func registerUserRouter(mux *http.ServeMux) {
	mux.HandleFunc("/fetch/", utils.CORS(controllers.GetUser))
	mux.HandleFunc("/update/", utils.CORS(controllers.UpdateUser))
}
