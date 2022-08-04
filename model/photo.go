package model

import "time"

type Photo struct {
	Id         int       `json:"Id"`
	Title      string    `json:"Title" validate:"required"`
	Caption    string    `json:"Caption"`
	Photo_Url  string    `json:"Photo_url" validate:"required"`
	User_Id    int       `json:"User_Id"`
	Created_at time.Time `json:"Created_at"`
	Updated_at time.Time `json:"Updated_at"`
}

type PhotoPostResponse struct {
	Id         int       `json:"Id"`
	Title      string    `json:"Title"`
	Caption    string    `json:"Caption"`
	Photo_Url  string    `json:"Photo_url"`
	User_Id    int       `json:"User_Id"`
	Created_at time.Time `json:"Created_at"`
}

type PhotoGetResponse struct {
	Photo `json:"Photo"`
	User  struct {
		Email    string `json:"Email"`
		Username string `json:"Username"`
	} `json:"User"`
}

type PhotoPutResponse struct {
	Id         int       `json:"Id"`
	Title      string    `json:"Title"`
	Caption    string    `json:"Caption"`
	Photo_Url  string    `json:"Photo_url"`
	User_Id    int       `json:"User_Id"`
	Updated_at time.Time `json:"Updated_at"`
}
