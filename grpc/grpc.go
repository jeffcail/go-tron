package grpc

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	_const "github.com/jeffcail/go-tron/common/const"

	"google.golang.org/grpc"

	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
)

// GetBowBlock
func (c *Client) GetBowBlock() (*api.BlockExtention, error) {
	err := c.keepConnect()
	if err != nil {
		return nil, err
	}
	block, err := c.GRPC.GetNowBlock()
	if err != nil {
		log.Fatal(err)
	}
	return block, nil
}

// GetBlockByNum
func (c *Client) GetBlockByNum(number int64) ([]*api.TransactionExtention, error) {
	err := c.keepConnect()
	if err != nil {
		return nil, err
	}
	block, err := c.GRPC.GetBlockByNum(number)
	if err != nil {
		return nil, err
	}
	return block.Transactions, nil
}

// GetTrxBalance
func (c *Client) GetTrxBalance(address string) (int64, error) {
	err := c.keepConnect()
	if err != nil {
		return int64(0), err
	}
	account, err := c.GRPC.GetAccount(address)
	if err != nil {
		return int64(0), err
	}
	return account.GetBalance(), nil
}

// GetTrc20Symbol
func (c *Client) GetTrc20Symbol(contractAddress string) (string, error) {
	err := c.keepConnect()
	if err != nil {
		return "", err
	}
	symbol, err := c.GRPC.TRC20GetSymbol(contractAddress)
	if err != nil {
		return "", err
	}
	return symbol, nil
}

// GetTrc20Name
func (c *Client) GetTrc20Name(contractAddress string) (string, error) {
	err := c.keepConnect()
	if err != nil {
		return "", err
	}
	name, err := c.GRPC.TRC20GetName(contractAddress)
	if err != nil {
		return "", err
	}
	return name, nil
}

// GetTrc20Decimal
func (c *Client) GetTrc20Decimal(contractAddress string) (*big.Int, error) {
	err := c.keepConnect()
	if err != nil {
		return big.NewInt(0), err
	}
	decimal, err := c.GRPC.TRC20GetDecimals(contractAddress)
	if err != nil {
		return big.NewInt(0), err
	}
	return decimal, nil
}

type Client struct {
	node string
	GRPC *client.GrpcClient
}

// NewClient
func NewClient(node string) (*Client, error) {
	c := new(Client)
	c.node = node
	c.GRPC = client.NewGrpcClient(node)

	opts := make([]grpc.DialOption, 0)
	opts = append(opts, grpc.WithInsecure())

	err := c.GRPC.Start(opts...)
	if err != nil {
		return nil, fmt.Errorf("grpc client start error: %v", err)
	}

	c.GRPC.SetAPIKey(_const.ApiKey)

	return c, nil
}

func (c *Client) SetTimeout(timeout time.Duration) error {
	if c == nil {
		return errors.New("client is nit ptr")
	}
	c.GRPC = client.NewGrpcClientWithTimeout(c.node, timeout)
	err := c.GRPC.Start()
	if err != nil {
		return fmt.Errorf("grpc start error: %v", err)
	}
	return nil
}

func (c *Client) keepConnect() error {
	_, err := c.GRPC.GetNodeInfo()
	if err != nil {
		if strings.Contains(err.Error(), "no such host") {
			return c.GRPC.Reconnect(c.node)
		}
		return fmt.Errorf("node connect error: %v", err)
	}
	return nil
}
