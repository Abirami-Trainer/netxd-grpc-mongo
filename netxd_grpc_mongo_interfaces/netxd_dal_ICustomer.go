package netxd_grpc_mongo_interfaces

import (
	"github.com/Abirami-Trainer/netxd-grpc-mongo/netxd_grpc_mongo_models"
)

type ICustomer interface {
	CreateCustomer(customer *netxd_grpc_mongo_models.Customer) (*netxd_grpc_mongo_models.CustomerResponse, error)
}
