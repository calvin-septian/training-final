package services

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"training-final/database"
	"training-final/model"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"

	jwt "github.com/golang-jwt/jwt/v4"
)

func PhotosHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	switch r.Method {
	case http.MethodPost:
		photoRegister(w, r)
	case http.MethodGet:
		getPhoto(w, r)
	case http.MethodPut:
		updatePhoto(w, r, id)
	case http.MethodDelete:
		deletePhoto(w, r, id)
	}
}

func photoRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)

	var photo model.Photo
	err := json.NewDecoder(r.Body).Decode(&photo)
	if err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	validate := validator.New()
	if err := validate.Struct(photo); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	photoId := database.DbConn.AddPhoto(photo, userInfo["Id"].(int), context.Background())
	photo = database.DbConn.GetPhotoById(photoId, context.Background())

	value := model.PhotoPostResponse{
		Id:         photo.Id,
		Title:      photo.Title,
		Caption:    photo.Caption,
		Photo_Url:  photo.Photo_Url,
		User_Id:    photo.User_Id,
		Created_at: photo.Created_at,
	}

	result, _ := json.Marshal(value)
	w.Write(result)
}

func getPhoto(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var photo model.Photo
	err := json.NewDecoder(r.Body).Decode(&photo)
	if err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	value := database.DbConn.GetAllPhoto(context.Background())

	result, _ := json.Marshal(value)
	w.Write(result)
}

func updatePhoto(w http.ResponseWriter, r *http.Request, photoId string) {
	w.Header().Add("Content-Type", "application/json")

	var photo model.Photo
	err := json.NewDecoder(r.Body).Decode(&photo)
	if err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	validate := validator.New()
	if err := validate.Struct(photo); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	photoid, _ := strconv.Atoi(photoId)
	database.DbConn.UpdatePhoto(photo, photoid, context.Background())
	photo = database.DbConn.GetPhotoById(photoid, context.Background())

	value := model.PhotoPutResponse{
		Id:         photo.Id,
		Title:      photo.Title,
		Caption:    photo.Caption,
		Photo_Url:  photo.Photo_Url,
		User_Id:    photo.User_Id,
		Updated_at: photo.Updated_at,
	}

	result, _ := json.Marshal(value)
	w.Write(result)
}

func deletePhoto(w http.ResponseWriter, r *http.Request, photoId string) {
	w.Header().Add("Content-Type", "application/json")
	photoid, _ := strconv.Atoi(photoId)
	database.DbConn.DeletePhoto(photoid, context.Background())

	type tmp struct {
		Message string `json:"message"`
	}
	value := tmp{Message: "Your photo has been successfully deleted"}
	result, _ := json.Marshal(value)

	w.Write(result)
}
