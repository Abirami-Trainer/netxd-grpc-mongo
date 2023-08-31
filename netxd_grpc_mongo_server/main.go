package netxd_grpc_mongo_server

import (
	"context"
	"fmt"
	"net"

	"github.com/Abirami-Trainer/netxd-grpc-mongo/netxd_grpc_mongo_config"
	"github.com/Abirami-Trainer/netxd-grpc-mongo/netxd_grpc_mongo_constants"
	"github.com/Abirami-Trainer/netxd-grpc-mongo/netxd_grpc_mongo_controllers"
	Netxd_Customer_git "github.com/Abirami-Trainer/netxd-grpc-mongo/netxd_grpc_mongo_proto"
	"github.com/Abirami-Trainer/netxd-grpc-mongo/netxd_grpc_mongo_services"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func InitDataBase(client *mongo.Client) {
	customerCollection := netxd_grpc_mongo_config.GetCollection(client, "Banking", "customer")
	netxd_grpc_mongo_controllers.CustomerService = netxd_grpc_mongo_services.InitCustomerService(customerCollection, context.Background())

}

func main() {
	mongoclient, err := netxd_grpc_mongo_config.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	InitDataBase(mongoclient)
	lis, err := net.Listen("tcp", netxd_grpc_mongo_constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	Netxd_Customer_git.RegisterCustomerServiceServer(s, &netxd_grpc_mongo_controllers.RPCServer{})
	fmt.Println("Server is listening on port ", netxd_grpc_mongo_constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
