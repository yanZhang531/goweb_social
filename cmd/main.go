package main

import (
	"github.com/pzlymformeet/social/pkg/routers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	routers.RegisterAllRouter(mux)
	http.ListenAndServe(":8010", mux)

}
