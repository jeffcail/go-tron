package grpc

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"

	"github.com/jeffcail/go-tron/common/sign"

	"github.com/jeffcail/go-tron/utils"

	_const "github.com/jeffcail/go-tron/common/const"

	"google.golang.org/grpc"

	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
)

// GetBowBlock
func (c *Client) GetBowBlock() (int64, error) {
	err := c.keepConnect()
	if err != nil {
		return 0, err
	}
	block, err := c.GRPC.GetNowBlock()
	if err != nil {
		log.Fatal(err)
	}
	return block.GetBlockHeader().RawData.Number, nil
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

// GetTrc10Token
func (c *Client) GetTrc10Token(assetID string) (string, error) {
	err := c.keepConnect()
	if err != nil {
		return "", err
	}
	asset, err := c.GRPC.GetAssetIssueByID(assetID)
	if err != nil {
		return "", err
	}
	return string(asset.Abbr), nil
}

// GetAssetIssueList
// 获取TRC10 通证列表
func (c *Client) GetTrc10TokenList(page int64, limit int) ([]*GetTrc10TokenListOut, error) {
	err := c.keepConnect()
	if err != nil {
		return nil, err
	}
	list, err := c.GRPC.GetAssetIssueList(page, limit)
	if err != nil {
		return nil, err
	}
	os := make([]*GetTrc10TokenListOut, 0)
	for _, v := range list.AssetIssue {
		o := &GetTrc10TokenListOut{
			ID:           v.Id,
			OwnerAddress: utils.HexToBase58(hex.EncodeToString(v.OwnerAddress)),
			Name:         string(v.Name),
			Abbr:         string(v.Abbr),
			Decimal:      v.Precision,
		}
		os = append(os, o)
	}
	return os, nil
}

// TransferTrx
func (c *Client) TransferTrx(from, to, pri string, amount int64) error {
	var (
		tx  *api.TransactionExtention
		err error
	)
	err = c.keepConnect()
	if err != nil {
		return err
	}

	tx, err = c.GRPC.Transfer(from, to, amount)
	if err != nil {
		return err
	}
	if tx == nil {
		return errors.New("transfer is nil")
	}

	signTx, err := sign.SignTransaction(tx.Transaction, pri)
	if err != nil {
		return err
	}
	if signTx == nil {
		return errors.New("after sign signTx is nil")
	}

	err = c.board(signTx)
	if err != nil {
		return err
	}
	return nil
}

// TransferTrc20
func (c *Client) TransferTrc20(from, pri, to, contractAddress string, amount, freeLimit int64) error {
	var (
		trc20Tx *api.TransactionExtention
		err     error
	)
	err = c.keepConnect()
	if err != nil {
		return err
	}
	a := big.NewInt(amount)
	trc20Tx, err = c.GRPC.TRC20Send(from, to, contractAddress, a, freeLimit)
	if err != nil {
		return err
	}
	if trc20Tx == nil {
		return errors.New("TRC20Send is nil")
	}
	signTrc20Tx, err := sign.SignTransaction(trc20Tx.Transaction, pri)
	if err != nil {
		return err
	}
	if signTrc20Tx == nil {
		return errors.New("after sign signTrc20Tx is nil")
	}
	err = c.board(signTrc20Tx)
	if err != nil {
		return err
	}
	return nil
}

// Board
func (c *Client) board(signTx *core.Transaction) error {
	err := c.keepConnect()
	if err != nil {
		return err
	}
	rs, err := c.GRPC.Broadcast(signTx)
	if err != nil {
		return fmt.Errorf("broadcast transaction error: %v", err)
	}
	if rs.Code != 0 {
		return fmt.Errorf("bad transaction: %v", string(rs.GetMessage()))
	}
	if rs.Result == true {
		return nil
	}
	d, _ := json.Marshal(rs)
	return fmt.Errorf("tx send fail: %s", string(d))
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
