package model

import "time"

type Comment struct {
	Id         int       `json:"Id"`
	User_Id    int       `json:"User_Id"`
	Photo_Id   int       `json:"Photo_Id"`
	Message    string    `json:"Message" validate:"required"`
	Created_at time.Time `json:"Created_at"`
	Updated_at time.Time `json:"Updated_at"`
}

type CommentPostResponse struct {
	Id         int       `json:"Id"`
	Message    string    `json:"Message"`
	Photo_Id   int       `json:"Photo_Id"`
	User_Id    int       `json:"User_Id"`
	Created_at time.Time `json:"Created_at"`
}

type CommentGetResponse struct {
	Comment `json:"Comment"`
	User    struct {
		Id       int    `json:"Id"`
		Email    string `json:"Email"`
		Username string `json:"Username"`
	} `json:"User"`
	Photo struct {
		Id        int    `json:"Id"`
		Title     string `json:"Title"`
		Caption   string `json:"Caption"`
		Photo_Url string `json:"Photo_url"`
		User_Id   int    `json:"User_Id"`
	} `json:"Photo"`
}

type CommentPutResponse struct {
	Id         int       `json:"Id"`
	Message    string    `json:"Message"`
	Photo_Id   int       `json:"Photo_Id"`
	User_Id    int       `json:"User_Id"`
	Updated_at time.Time `json:"Updated_at"`
}
