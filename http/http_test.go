package http

import (
	"fmt"
	"testing"

	"github.com/jeffcail/go-tron/utils"
)

func TestGetNowBlock(t *testing.T) {
	height, err := GetNowBlock()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("block height: ", height)
}

func TestGetBlockByNum(t *testing.T) {
	num := 46798578
	res, err := GetBlockByNum(int64(num))
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range res.Transactions {
		fmt.Println(v)
	}
}

func TestGetTrxBalance(t *testing.T) {
	address := "TCyps3Pber1ghKYgw5vq6KLxFVPJk9EvWC"
	balance, err := GetTrxBalance(address)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(balance)
}

func TestGetTrc10Balance(t *testing.T) {
	address := "TK1UXQBkvAwBypz1bTWcuLHFaB8JmTjoUw"
	asssetId := "1002000"
	hex, _ := utils.Base58ToHex(address)
	h := hex[2:]
	balance, err := GetTrc10Balance(h, asssetId)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(balance)
}

func TestGetTrc20Symbol(t *testing.T) {
	contract := "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
	hex, _ := utils.Base58ToHex(contract)
	h := hex[2:]
	symbol, err := GetTrc20Symbol(h)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(symbol)
}

func TestGetTrc10Token(t *testing.T) {
	assetID := "1000001"
	trc10Token, err := GetTrc10Token(assetID)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(trc10Token)
}
