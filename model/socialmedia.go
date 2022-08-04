package model

import "time"

type SocialMedia struct {
	Id               int       `json:"Id"`
	Name             string    `json:"Name" validate:"required"`
	Social_media_url string    `json:"Social_media_url" validate:"required"`
	User_Id          int       `json:"UserId"`
	Created_at       time.Time `json:"Created_at"`
	Updated_at       time.Time `json:"Updated_at"`
}

type SocialMediaPostResponse struct {
	Id               int       `json:"Id"`
	Name             string    `json:"Name"`
	Social_media_url string    `json:"Social_media_url"`
	User_Id          int       `json:"UserId"`
	Created_at       time.Time `json:"Created_at"`
}

type SocialMediaGetResponse struct {
	SocialMedia `json:"SocialMedia"`
	User        struct {
		Id       int    `json:"Id"`
		Username string `json:"Username"`
	} `json:"User"`
}

type SocialMediaPutResponse struct {
	Id               int       `json:"Id"`
	Name             string    `json:"Name"`
	Social_media_url string    `json:"Social_media_url"`
	User_Id          int       `json:"UserId"`
	Updated_at       time.Time `json:"Updated_at"`
}
