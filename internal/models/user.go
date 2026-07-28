package models

import (
	"time"
)

// define 3 roles: admin, user, guest
var (
	RoleAdmin = "admin"
	RoleUser  = "user"
	RoleGuest = "guest"
)

type User struct {
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	FirstName string    `bson:"firstname" json:"firstname" required:"true"`
	LastName  string    `bson:"lastname" json:"lastname"`
	Role      string    `bson:"role" json:"role" default:"user"`
	Email     string    `bson:"email" json:"email" required:"true"`
	UserId    string    `bson:"user_id" json:"user_id"`
	Username  string    `bson:"username" json:"username" required:"true"`
	Password  string    `bson:"password" json:"password" required:"true"`
}
