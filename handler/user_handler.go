package handler

import (
	"encoding/json"
	"fmt"
	"fossil/api"
	"fossil/model"
	"fossil/service"
	"log"
	"net/http"
)

func UserList(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("route:user list"))
}

func UserDetail(rw http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	if username == "" {
		rw.Write([]byte("route:user detail | broken url"))
	}

	rw.Header().Set("Content-Type", "application/json")

	json.NewEncoder(rw).Encode(model.User{
		Username: username,
	})
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {

	userRequest := new(api.CreateUserRequest)
	result := new(model.User)

	err := api.ParseRequestJSON(rw, r, userRequest)
	if err != nil {
		api.JSONResponse(rw, http.StatusBadRequest, err.Error(), nil)
		return
	}

	log.Println(*userRequest)
	ok, validationError := api.ValidateStruct(userRequest)

	if !ok {
		api.JSONResponse(rw, http.StatusBadRequest, "Validation error", validationError)
		return
	}

	err = service.UserServiceInstance.AddUser(*userRequest, result)
	if err != nil {
		api.JSONResponse(rw, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// success, returns created user
	api.JSONResponse(rw, http.StatusOK, "User created", *result)
	return
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	if username == "" {
		rw.Write([]byte("route:user update | broken url"))
	}
	rw.Write([]byte(fmt.Sprintf("route:user update (%s) ", username)))
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	if username == "" {
		rw.Write([]byte("route:user delete | broken url"))
	}
	rw.Write([]byte(fmt.Sprintf("route:user delete (%s) ", username)))
}
