package http

import (
	"fmt"
	"testing"
)

func TestGetTrc10Token(t *testing.T) {
	assetID := "1000001"
	trc10Token, err := GetTrc10Token(assetID)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(trc10Token)
}
