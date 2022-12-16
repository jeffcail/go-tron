package http

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	_const "github.com/jeffcail/go-tron/common/const"

	"github.com/jeffcail/go-tron/common/req"
)

// GetNowBlock
func GetNowBlock() (int64, error) {
	url := fmt.Sprintf("%s%s", _const.HttpApi, "/wallet/getnowblock")
	h := buildHeader()
	bytes, err := req.Get(url, h, nil)
	if err != nil {
		return 0, err
	}
	o := &GetNowBlockOut{}
	err = json.Unmarshal(bytes, o)
	if err != nil {
		return 0, err
	}
	return o.BlockHeader.RawData.Number, nil
}

// GetTrc10Token
func GetTrc10Token(assetID string) (string, error) {
	url := fmt.Sprintf("%s%s", _const.HttpApi, "wallet/getassetissuebyid")
	h := buildHeader()
	p := make(map[string]interface{})
	p["value"] = assetID
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

func buildHeader() map[string]string {
	header := make(map[string]string)
	header["accept"] = "application/json"
	header["content-type"] = "application/json"
	header["x-api-key"] = _const.ApiKey
	return header
}
