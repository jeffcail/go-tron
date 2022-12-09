package utils

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

// HexToBigInt
func HexToBigInt(hex string) *big.Int {
	if hex == "" {
		hex = "0x0"
	}
	hex = strings.Replace(hex, "0x0", "", -1)
	n := new(big.Int)
	n, _ = n.SetString(hex, 16)
	return n
}

// Int64ToHex
func Int64ToHex(a int64) string {
	b := strconv.FormatInt(a, 16)
	return fmt.Sprintf("%s%s", "0x", b)
}
