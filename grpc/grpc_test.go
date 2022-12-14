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
