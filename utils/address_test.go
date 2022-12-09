package utils

import (
	"fmt"
	"testing"
)

func TestBase58ToHex(t *testing.T) {
	b58 := "THp5sWs7WzHxF3U7ytBVLqzPAR5iFWey8D"
	hex, err := Base58ToHex(b58)
	if err != nil {
		t.Fatal(hex)
	}
	fmt.Println(hex)
}

func TestHexToBase58(t *testing.T) {
	hex := "4156074F8CD2F65BB17788F047DA52B5A5CB372298"
	base58 := HexToBase58(hex)
	fmt.Println(base58)
}
