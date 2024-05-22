package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	Name          string             `json:"name"`
	Email         string             `json:"email"`
	ID            primitive.ObjectID `bson:"_id"`
	RollNo        string             `json:"rollno"`
	Password      *string            `json:"password"`
	Token         *string            `json:"token"`
	Refresh_token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	Image         *string            `json:"image"`
}
type Attendance struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	UserEmail string             `json:"email"`
	Date      string             `json:"date"`
	Status    string             `bson:"status" json:"status"`
	RollNo    string             `json:"rollno`
}
type Admin struct {
	Email         string             `json:"email"`
	ID            primitive.ObjectID `bson:"_id"`
	RollNo        int                `json:"rollno"`
	Password      *string            `json:"password"`
	Token         *string            `json:"token"`
	Refresh_token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
}
