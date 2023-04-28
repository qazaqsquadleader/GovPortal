package models

import "time"

type User struct {
	UserId        int       `json:userId`
	Username      string    `json:username`
	Password      string    `json:userPassword`
	Email         string    `json:userEmail`
	Token         string    `json:userToken`
	TokenDuration time.Time `json:userTokenCreate`
	Status        string    `json:userStatus`
}
