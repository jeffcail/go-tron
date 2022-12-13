package http

import (
	"fmt"

	"github.com/jeffcail/go-tron/common/req"
)

var (
	api = "https://api.trongrid.io"
	key = "c1f9b888-62e7-4f9e-8c80-f2aad8313feb"
)

// 获取最新区块高度
func GetNowBlock(header map[string]string) {
	url := fmt.Sprintf("%s%s", api, "/wallet/getnowblock")
	req.Get(url, header, nil)
}
