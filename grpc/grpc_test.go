package grpc

import (
	"fmt"
	"log"
	"testing"
)

var node = "52.53.189.99:50051"

func TestGetBowBlock(t *testing.T) {
	c, err := NewClient(node)
	if err != nil {
		log.Fatal(err)
	}
	block, err := c.GetBowBlock()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(block.GetBlockHeader().RawData.Number)
}

func TestClient_GetBlockByNum(t *testing.T) {
	number := 46798578
	c, err := NewClient(node)
	if err != nil {
		log.Fatal(err)
	}
	transactions, err := c.GetBlockByNum(int64(number))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(transactions[0].Transaction.RawData.Contract[0].Type)
	fmt.Println(len(transactions))
	fmt.Println(transactions[0])
}
