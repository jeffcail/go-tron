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

func TestClient_GetTrxBalance(t *testing.T) {
	c, err := NewClient(node)
	if err != nil {
		log.Fatal(err)
	}
	address := "TDMDMXnFpkqrBEVCjEHiwRHZ6UQroe2j74"
	balance, err := c.GetTrxBalance(address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)
}

func TestClient_GetTrc20Symbol(t *testing.T) {
	c, err := NewClient(node)
	if err != nil {
		log.Fatal(err)
	}
	address := "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
	symbol, err := c.GetTrc20Symbol(address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(symbol)
}

func TestClient_GetTrc20Name(t *testing.T) {
	c, err := NewClient(node)
	if err != nil {
		log.Fatal(err)
	}
	address := "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
	name, err := c.GetTrc20Name(address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}

func TestClient_GetTrc20Decimal(t *testing.T) {
	c, err := NewClient(node)
	if err != nil {
		log.Fatal(err)
	}
	address := "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
	decimal, err := c.GetTrc20Decimal(address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(decimal)
}
