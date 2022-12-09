package generate

import (
	"encoding/hex"

	"github.com/btcsuite/btcd/btcec"
	addr "github.com/fbsobreira/gotron-sdk/pkg/address"
)

// 离线生成波场地址和密钥
func GenerateKey() (wif string, address string) {
	pri, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return "", ""
	}
	if len(pri.D.Bytes()) != 32 {
		for {
			pri, err = btcec.NewPrivateKey(btcec.S256())
			if err != nil {
				continue
			}
			if len(pri.D.Bytes()) == 32 {
				break
			}
		}
	}
	address = addr.PubkeyToAddress(pri.ToECDSA().PublicKey).String()
	wif = hex.EncodeToString(pri.D.Bytes())
	return
}
