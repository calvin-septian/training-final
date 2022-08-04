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

func SocialMediasHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	switch r.Method {
	case http.MethodPost:
		socialMediaRegister(w, r)
	case http.MethodGet:
		getSocialMedia(w, r)
	case http.MethodPut:
		updateSocialMedia(w, r, id)
	case http.MethodDelete:
		deleteSocialMedia(w, r, id)
	}
}

func socialMediaRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)

	var sosmed model.SocialMedia
	err := json.NewDecoder(r.Body).Decode(&sosmed)
	if err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	validate := validator.New()
	if err := validate.Struct(sosmed); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	sosmedId := database.DbConn.AddSocialMedia(sosmed, userInfo["Id"].(int), context.Background())
	sosmed = database.DbConn.GetSocialMediaById(sosmedId, context.Background())

	value := model.SocialMediaPostResponse{
		Id:               sosmed.Id,
		Name:             sosmed.Name,
		Social_media_url: sosmed.Social_media_url,
		User_Id:          sosmed.User_Id,
		Created_at:       sosmed.Created_at,
	}

	result, _ := json.Marshal(value)
	w.Write(result)
}

func getSocialMedia(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var sosmed model.SocialMedia
	err := json.NewDecoder(r.Body).Decode(&sosmed)
	if err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	value := database.DbConn.GetAllSocialMedia(context.Background())

	type tmp struct {
		Social_Medias []model.SocialMediaGetResponse `json:"Social_Medias"`
	}
	tmpValue := tmp{Social_Medias: value}
	result, _ := json.Marshal(tmpValue)
	w.Write(result)
}

func updateSocialMedia(w http.ResponseWriter, r *http.Request, sosmedId string) {
	w.Header().Add("Content-Type", "application/json")

	var sosmed model.SocialMedia
	err := json.NewDecoder(r.Body).Decode(&sosmed)
	if err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	validate := validator.New()
	if err := validate.Struct(sosmed); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	sosmedid, _ := strconv.Atoi(sosmedId)
	database.DbConn.UpdateSocialMedia(sosmed, sosmedid, context.Background())
	sosmed = database.DbConn.GetSocialMediaById(sosmedid, context.Background())

	value := model.SocialMediaPutResponse{
		Id:               sosmed.Id,
		Name:             sosmed.Name,
		Social_media_url: sosmed.Social_media_url,
		User_Id:          sosmed.User_Id,
		Updated_at:       sosmed.Updated_at,
	}

	result, _ := json.Marshal(value)
	w.Write(result)
}

func deleteSocialMedia(w http.ResponseWriter, r *http.Request, sosmedId string) {
	w.Header().Add("Content-Type", "application/json")
	sosmedid, _ := strconv.Atoi(sosmedId)
	database.DbConn.DeleteComment(sosmedid, context.Background())

	type tmp struct {
		Message string `json:"Message"`
	}
	value := tmp{Message: "Your social media has been successfully deleted"}
	result, _ := json.Marshal(value)

	w.Write(result)
}
