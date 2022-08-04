package model

import "time"

type User struct {
	Id         int       `json:"Id"`
	Username   string    `json:"Username" validate:"required"`
	Email      string    `json:"Email" validate:"required,email"`
	Password   string    `json:"Password" validate:"required,min=6"`
	Age        int       `json:"Age" validate:"required,gte=8"`
	Created_at time.Time `json:"Created_at"`
	Updated_at time.Time `json:"Updated_at"`
}

type UserPostResponse struct {
	Id         int       `json:"Id"`
	Username   string    `json:"Username"`
	Email      string    `json:"Email"`
	Age        int       `json:"Age"`
	Created_at time.Time `json:"Created_at"`
}

type UserPutResponse struct {
	Id         int       `json:"Id"`
	Username   string    `json:"Username"`
	Email      string    `json:"Email"`
	Age        int       `json:"Age"`
	Updated_at time.Time `json:"Updated_at"`
}
