package netxd_grpc_mongo_controllers

import (
	"context"

	cus "github.com/Abirami-Trainer/netxd-grpc-mongo/netxd_grpc_mongo_proto"

	"github.com/Abirami-Trainer/netxd-grpc-mongo/netxd_grpc_mongo_models"

	"github.com/Abirami-Trainer/netxd-grpc-mongo/netxd_grpc_mongo_interfaces"
)

type RPCServer struct {
	cus.UnimplementedCustomerServiceServer
}

var (
	CustomerService netxd_grpc_mongo_interfaces.ICustomer
)

func (s *RPCServer) CreateCustomer(ctx context.Context, req *cus.Customer) (*cus.CustomerResponse, error) {
	dbCustomer := &netxd_grpc_mongo_models.Customer{FirstName: req.FirstName, LastName: req.LastName}
	result, err := CustomerService.CreateCustomer(dbCustomer)
	if err != nil {
		return nil, err
	} else {
		responseCustomer := &cus.CustomerResponse{
			CustomerId: result.CustomerId,
		}
		return responseCustomer, nil
	}
}
