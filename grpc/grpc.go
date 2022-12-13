package grpc

import (
	"log"
	"time"

	"github.com/fbsobreira/gotron-sdk/pkg/client"
)

var (
	node = "3.225.171.164:50051"
	Conn *Client
)

type Client struct {
	node string
	grpc *client.GrpcClient
}

func init() {
	c := new(Client)
	c.node = node
	c.grpc = client.NewGrpcClientWithTimeout(node, 10*time.Second)
	err := c.grpc.Start()
	if err != nil {
		log.Fatal(err)
	}
	Conn = c
}
