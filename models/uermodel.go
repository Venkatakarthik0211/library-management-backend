package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    *string            `json:"first_name" validate:"required, min = 2, max = 100"`
	Last_name     *string            `json:"first_name" validate:"required, min = 2, max = 100"`
	Password      *string            `json:"password" validate:"required, min = 6"`
	Email         *string            `jsom:"email" validate:"email, required"`
	Phone         *string            `json:"phone" validate:"required"`
	Token         *string            `json:"token"`
	User_type     *string            `json:"user_type" validate:"required, eq=ADMIN|eq=USER"`
	Refresh_token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	UPdated_at    time.Time          `json:"updated_at"`
	ISBN          *string            `json:"isbn"`
	Title         *string            `json:"title"`
	Author        *string            `json:"author"`
	Price         *string            `json:"price"`
	Quantity      *string            `json:"quantity"`
}
