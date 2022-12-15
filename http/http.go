package http

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	_const "github.com/jeffcail/go-tron/common/const"

	"github.com/jeffcail/go-tron/common/req"
)

var (
	api = "https://api.trongrid.io"
)

// 获取最新区块高度
func GetNowBlock(header map[string]string) {
	url := fmt.Sprintf("%s%s", api, "/wallet/getnowblock")
	req.Get(url, header, nil)
}

// GetTrc10Token
func GetTrc10Token(assetID string) (string, error) {
	url := fmt.Sprintf("%s%s", _const.HttpApi, "wallet/getassetissuebyid")
	h, p := buildIdentifyTransactionToken(assetID)
	res, err := req.Post(url, h, p)
	if err != nil {
		return "", errors.New(fmt.Sprintf("TRC10 获取币种失败 err: %v", err))
	}
	asset := &GetAssetIssueByID{}
	_ = json.Unmarshal(res, &asset)
	token, err := hex.DecodeString(asset.Abbr)
	if err != nil {
		return "", errors.New(fmt.Sprintf("TRC10 解析币种失败 err: %v", err))
	}
	return string(token), nil
}

func buildIdentifyTransactionToken(hash string) (map[string]string, map[string]interface{}) {
	header := make(map[string]string)
	header["accept"] = "application/json"
	header["content-type"] = "application/json"
	header["x-api-key"] = _const.ApiKey

	p := make(map[string]interface{})
	p["value"] = hash
	return header, p
}
