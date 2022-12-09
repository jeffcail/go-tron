package utils

import (
	"fmt"
	"strings"

	"github.com/fbsobreira/gotron-sdk/pkg/address"
)

// Base58ToHex
func Base58ToHex(b58Address string) (string, error) {
	a, err := address.Base58ToAddress(b58Address)
	if err != nil {
		return "", err
	}
	return strings.ToUpper(a.Hex()), nil
}

// HexToBase58
func HexToBase58(hexAddress string) string {
	hexAddress = fmt.Sprintf("%s%s", "0x", hexAddress)
	a := address.HexToAddress(hexAddress)
	return a.String()
}
