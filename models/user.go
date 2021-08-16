package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserReq struct {
	Email string `bson:"email" json:"email"`
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserReq   `json:",inline" bson:",inline"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}
