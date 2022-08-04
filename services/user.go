package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"training-final/database"
	"training-final/helper"
	"training-final/model"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	switch r.Method {
	case http.MethodPost:
		if r.URL.Path == "register" {
			userRegister(w, r)
		} else {
			userLogin(w, r)
		}
	case http.MethodPut:
		updateUser(w, r, id)
	case http.MethodDelete:
		deleteUser(w, r, id)
	}
}

func userRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte("error decoding json body"))
		return
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	pass := []byte(user.Password)

	bcrypthash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	user.Password = string(bcrypthash)

	userId := database.DbConn.AddUser(user, context.Background())
	user = database.DbConn.GetUserById(userId, context.Background())

	value := model.UserPostResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age}

	result, _ := json.Marshal(value)
	w.Write(result)
}

func userLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	type dataLogin struct {
		Email    string `json:"Email" validate:"required,email"`
		Password string `json:"Password" validate:"required,min=6"`
	}
	var userlogin dataLogin
	err := json.NewDecoder(r.Body).Decode(&userlogin)
	if err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	validate := validator.New()
	if err := validate.Struct(userlogin); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	user := database.DbConn.GetPasswordByEmail(userlogin.Email, context.Background())

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userlogin.Password))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	tokenString, err := helper.GenerateJWT(user)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	result, _ := json.Marshal(tokenString)

	w.Write(result)
}

func updateUser(w http.ResponseWriter, r *http.Request, userId string) {
	w.Header().Add("Content-Type", "application/json")

	type dataUpdate struct {
		Email    string `json:"Email" validate:"required,email"`
		Username string `json:"Username" validate:"required"`
	}
	var dataupdate dataUpdate
	err := json.NewDecoder(r.Body).Decode(&dataupdate)
	if err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	validate := validator.New()
	if err := validate.Struct(dataupdate); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	user := model.User{Email: dataupdate.Email, Username: dataupdate.Username}

	userid, _ := strconv.Atoi(userId)
	database.DbConn.UpdateUser(user, userid, context.Background())
	user = database.DbConn.GetUserById(userid, context.Background())

	value := model.UserPutResponse{
		Id:         user.Id,
		Username:   user.Username,
		Email:      user.Email,
		Age:        user.Age,
		Updated_at: user.Updated_at,
	}
	result, _ := json.Marshal(value)

	w.Write(result)
}

func deleteUser(w http.ResponseWriter, r *http.Request, userId string) {
	w.Header().Add("Content-Type", "application/json")
	userid, _ := strconv.Atoi(userId)
	database.DbConn.DeleteUser(userid, context.Background())

	type tmp struct {
		Message string `json:"Message"`
	}
	value := tmp{Message: "Your account has been successfully deleted"}
	result, _ := json.Marshal(value)

	w.Write(result)
}
