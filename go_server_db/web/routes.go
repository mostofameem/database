package web

import (
	"go_server_db/web/handlers"
	"go_server_db/web/middlewares"
	"net/http"
)

func InitRouts(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"POST /users",
		manager.With(
			http.HandlerFunc(handlers.Create),
		),
	)

	mux.Handle(
		"GET /users",
		manager.With(
			http.HandlerFunc(handlers.Read),
		),
	)

	mux.Handle(
		"DELETE /users/{id}",
		manager.With(
			http.HandlerFunc(handlers.Delete),
		),
	)

	mux.Handle(
		"PUT /users/{id}",
		manager.With(
			http.HandlerFunc(handlers.Update),
		),
	)
}
