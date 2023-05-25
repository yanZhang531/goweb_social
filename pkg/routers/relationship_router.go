package routers

import (
	"github.com/pzlymformeet/social/pkg/controllers"
	"github.com/pzlymformeet/social/pkg/utils"
	"net/http"
)

var registerRelationship = func(mux *http.ServeMux) {
	mux.HandleFunc("/relationship/", utils.CORS(controllers.DispatchRelationship))
}
