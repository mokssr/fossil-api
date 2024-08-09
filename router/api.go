package router

import (
	"fossil/handler"
	"net/http"
)

func MakeAPIHandler(prefix string) http.Handler {
	api := http.NewServeMux()

	api.HandleFunc("/", handleApiRoot)

	api.HandleFunc("GET /user", handler.UserList)
	api.HandleFunc("POST /user", handler.CreateUser)
	api.HandleFunc("GET /user/{username}", handler.UserDetail)
	api.HandleFunc("PUT /user/{username}", handler.UpdateUser)
	api.HandleFunc("DELETE /user/{username}", handler.DeleteUser)

	return http.StripPrefix(prefix, api)
}

func handleApiRoot(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("route not found"))
}
