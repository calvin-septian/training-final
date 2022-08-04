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

func CommentsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	switch r.Method {
	case http.MethodPost:
		commentRegister(w, r)
	case http.MethodGet:
		getComments(w, r)
	case http.MethodPut:
		updateComment(w, r, id)
	case http.MethodDelete:
		deleteComment(w, r, id)
	}
}

func commentRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)

	var comment model.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	validate := validator.New()
	if err := validate.Struct(comment); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	commentId := database.DbConn.AddComment(comment, userInfo["Id"].(int), context.Background())
	comment = database.DbConn.GetCommentById(commentId, context.Background())

	value := model.CommentPostResponse{
		Id:         comment.Id,
		Message:    comment.Message,
		Photo_Id:   comment.Photo_Id,
		User_Id:    comment.User_Id,
		Created_at: comment.Created_at,
	}

	result, _ := json.Marshal(value)
	w.Write(result)
}

func getComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var comment model.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	value := database.DbConn.GetAllComments(context.Background())

	result, _ := json.Marshal(value)
	w.Write(result)
}

func updateComment(w http.ResponseWriter, r *http.Request, commentId string) {
	w.Header().Add("Content-Type", "application/json")

	var comment model.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	validate := validator.New()
	if err := validate.Struct(comment); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	commentid, _ := strconv.Atoi(commentId)
	database.DbConn.UpdateComment(comment, commentid, context.Background())
	comment = database.DbConn.GetCommentById(commentid, context.Background())

	value := model.CommentPutResponse{
		Id:         comment.Id,
		Message:    comment.Message,
		Photo_Id:   comment.Photo_Id,
		User_Id:    comment.User_Id,
		Updated_at: comment.Updated_at,
	}

	result, _ := json.Marshal(value)
	w.Write(result)
}

func deleteComment(w http.ResponseWriter, r *http.Request, commentId string) {
	w.Header().Add("Content-Type", "application/json")
	commentid, _ := strconv.Atoi(commentId)
	database.DbConn.DeleteComment(commentid, context.Background())

	type tmp struct {
		Message string `json:"message"`
	}
	value := tmp{Message: "Your comment has been successfully deleted"}
	result, _ := json.Marshal(value)

	w.Write(result)
}
