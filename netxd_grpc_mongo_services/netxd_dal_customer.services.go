package netxd_grpc_mongo_services

import (
	"context"
	"errors"
	"time"

	"github.com/Abirami-Trainer/netxd-grpc-mongo/netxd_grpc_mongo_interfaces"
	"github.com/Abirami-Trainer/netxd-grpc-mongo/netxd_grpc_mongo_models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	CustomerCollection *mongo.Collection
	ctx                context.Context
}

// CreateCustomer implements netxd_dal_interfaces.ICustomer.
func (cus *CustomerService) CreateCustomer(customer *netxd_grpc_mongo_models.Customer) (*netxd_grpc_mongo_models.CustomerResponse, error) {
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = customer.CreatedAt
	customer.IsActive = true
	res, err := cus.CustomerCollection.InsertOne(cus.ctx, &customer)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 1100 {
			return nil, errors.New("customer with Id already exists")
		}
		return nil, err
	}
	var newCustomer *netxd_grpc_mongo_models.CustomerResponse
	query := bson.M{"id": res.InsertedID}
	err = cus.CustomerCollection.FindOne(cus.ctx, query).Decode(&newCustomer)
	if err != nil {
		return nil, err
	}
	return newCustomer, nil
}

func InitCustomerService(collection *mongo.Collection, ctx context.Context) netxd_grpc_mongo_interfaces.ICustomer {

	return &CustomerService{collection, ctx}
}
