package netxd_grpc_mongo_models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	CustomerId string    `json:"customer_id" bson:"customer_id"`
	FirstName  string    `json:"first_name" bson:"first_name"`
	LastName   string    `json:"last_name" bson:"last_name"`
	BankId     string    `json:"bank_id" bson:"bank_id"`
	Balance    float32   `json:"balance" bson:"balance"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
	IsActive   bool      `json:"is_active" bson:"is_active"`
}
type CustomerResponse struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	CustomerId string             `json:"customer_id" bson:"customer_id"`
	CreatedAt  string             `json:"created_at" bson:"created_at"`
}
