package utils

import (
	"math/big"
	"strings"
)

// ParseData
func ParseData(data string) (string, *big.Int) {
	if len(data) == 8 {
		return "", big.NewInt(0)
	}

	hexAddress := strings.TrimPrefix(data[8:72], "0000000000000000000000")
	hexAmount := strings.TrimPrefix(data[72:], "0000000000000000000000000000000000000000000000000000000")
	if hexAmount == "" {
		return hexAddress, big.NewInt(0)
	}
	amount := HexToBigInt(hexAmount)
	return hexAddress, amount
}
