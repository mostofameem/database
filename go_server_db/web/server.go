package web

import (
	"go_server_db/web/middlewares"
	"net/http"
)

func StartServer() *http.ServeMux {
	manager := middlewares.NewManager()
	mux := http.NewServeMux()
	InitRouts(mux, manager)
	return mux
}
