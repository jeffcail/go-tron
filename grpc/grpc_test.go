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
	height, err := c.GetBowBlock()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(height)
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

func TestClient_GetTrc10Token(t *testing.T) {
	c, err := NewClient(node)
	if err != nil {
		log.Fatal(err)
	}
	assetID := "1000001"
	trc10Token, err := c.GetTrc10Token(assetID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(trc10Token)
}

func TestClient_GetTrc10TokenList(t *testing.T) {
	c, err := NewClient(node)
	if err != nil {
		log.Fatal(err)
	}
	outs, err := c.GetTrc10TokenList(1, 300)
	fmt.Println(len(outs))
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range outs {
		fmt.Println(fmt.Sprintf("ID:【%s】 OwnerAddress:【%s】 Name:【%s】 Abbr:【%s】 Deciaml:【%v】",
			v.ID, v.OwnerAddress, v.Name, v.Abbr, v.Decimal))
	}
}

func TestClient_TransferTrx(t *testing.T) {
	c, err := NewClient(node)
	if err != nil {
		t.Fatal(err)
	}
	// 1 SUN
	pri := ""
	from := "TDMDMXnFpkqrBEVCjEHiwRHZ6UQroe2j74"
	to := "TCyps3Pber1ghKYgw5vq6KLxFVPJk9EvWC"
	amount := 1
	err = c.TransferTrx(from, to, pri, int64(amount))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("TRX 转账成功")
}

func TestClient_TransferTrc20(t *testing.T) {
	c, err := NewClient(node)
	if err != nil {
		t.Fatal(err)
	}
	// USDT
	pri := ""
	from := "TDMDMXnFpkqrBEVCjEHiwRHZ6UQroe2j74"
	to := "TCyps3Pber1ghKYgw5vq6KLxFVPJk9EvWC"
	contractAddress := "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
	amount := 1
	freeLimit := 50000000 // 50 TRX
	err = c.TransferTrc20(from, pri, to, contractAddress, int64(amount), int64(freeLimit))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("TRC20 转账成功")
}
