package model

import "time"

// import (
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

type User struct{
	ID string
	Username string
	Displayname string
	Email string
	Password string
	CreatedOn time.Time
	Twofactor bool
	OTP string
}
