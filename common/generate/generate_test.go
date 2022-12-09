package generate

import (
	"fmt"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	key, address := GenerateKey()
	fmt.Println(address)
	fmt.Println(key)
}
