package netxd_grpc_mongo_client

import (
	"context"
	"fmt"
	"log"

	Netxd_Customer_git "github.com/Abirami-Trainer/netxd-grpc-mongo/netxd_grpc_mongo_proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect :%v", err)

	}
	defer conn.Close()
	client := Netxd_Customer_git.NewCustomerServiceClient(conn)
	response, err := client.CreateCustomer(context.Background(), &Netxd_Customer_git.Customer{FirstName: "Arun", LastName: "kumar"})
	if err != nil {
		log.Fatal("Failed to call SayHello:%v", err)
	}
	fmt.Printf("Response: %s\n", response.CustomerId)
}
