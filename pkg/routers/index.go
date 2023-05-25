package routers

import "net/http"

var RegisterAllRouter = func(mux *http.ServeMux) {
	registerAuthRouter(mux)
	registerUserRouter(mux)
	registerRelationship(mux)
	regiseterPostRouter(mux)
	RegisterCommentRouter(mux)
	registerLikeRouter(mux)
}
