package http

// GetNowBlockOut
type GetNowBlockOut struct {
	BlockID     string `json:"blockID"`
	BlockHeader struct {
		RawData struct {
			Number         int64  `json:"number"`
			TxTrieRoot     string `json:"txTrieRoot"`
			WitnessAddress string `json:"witness_address"`
			ParentHash     string `json:"parentHash"`
			Version        int64  `json:"version"`
			Timestamp      int64  `json:"timestamp"`
		} `json:"raw_data"`
		WitnessSignature string `json:"witness_signature"`
	} `json:"block_header"`
}

// GetBlockByNumOut
type GetBlockByNumOut struct {
	BlockID     string `json:"blockID"`
	BlockHeader struct {
		RawData struct {
			Number         int    `json:"number"`
			TxTrieRoot     string `json:"txTrieRoot"`
			WitnessAddress string `json:"witness_address"`
			ParentHash     string `json:"parentHash"`
			Version        int    `json:"version"`
			Timestamp      int64  `json:"timestamp"`
		} `json:"raw_data"`
		WitnessSignature string `json:"witness_signature"`
	} `json:"block_header"`
	Transactions []struct {
		Ret []struct {
			ContractRet string `json:"contractRet"`
		} `json:"ret"`
		Signature []string `json:"signature"`
		TxID      string   `json:"txID"`
		RawData   struct {
			Contract []struct {
				Parameter struct {
					Value struct {
						Amount       int    `json:"amount"`
						OwnerAddress string `json:"owner_address"`
						ToAddress    string `json:"to_address"`
					} `json:"value"`
					TypeURL string `json:"type_url"`
				} `json:"parameter"`
				Type string `json:"type"`
			} `json:"contract"`
			RefBlockBytes string `json:"ref_block_bytes"`
			RefBlockHash  string `json:"ref_block_hash"`
			Expiration    int64  `json:"expiration"`
			Timestamp     int64  `json:"timestamp"`
		} `json:"raw_data,omitempty"`
		RawDataHex string `json:"raw_data_hex"`
	}
}

// GetTrxBalanceOut
type GetTrxBalanceOut struct {
	Data []struct {
		LatestOprationTime int64 `json:"latest_opration_time"`
		OwnerPermission    struct {
			Keys []struct {
				Address string `json:"address"`
				Weight  int    `json:"weight"`
			} `json:"keys"`
			Threshold      int    `json:"threshold"`
			PermissionName string `json:"permission_name"`
		} `json:"owner_permission"`
		AccountResource struct {
			LatestConsumeTimeForEnergy int64 `json:"latest_consume_time_for_energy"`
		} `json:"account_resource"`
		ActivePermission []struct {
			Operations string `json:"operations"`
			Keys       []struct {
				Address string `json:"address"`
				Weight  int    `json:"weight"`
			} `json:"keys"`
			Threshold      int    `json:"threshold"`
			ID             int    `json:"id"`
			Type           string `json:"type"`
			PermissionName string `json:"permission_name"`
		} `json:"active_permission"`
		Address    string `json:"address"`
		Balance    int64  `json:"balance"`
		CreateTime int64  `json:"create_time"`
		Trc20      []struct {
			TR7NHqjeKQxGTCi8Q8ZY4PL8OtSzgjLj6T string `json:"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"`
		} `json:"trc20"`
		LatestConsumeFreeTime int64 `json:"latest_consume_free_time"`
	} `json:"data"`
	Success bool `json:"success"`
	Meta    struct {
		At       int64 `json:"at"`
		PageSize int   `json:"page_size"`
	} `json:"meta"`
}

// GetTrc10BalanceOut
type GetTrc10BalanceOut struct {
	Address string `json:"address"`
	Balance int    `json:"balance"`
	Frozen  []struct {
		FrozenBalance int64 `json:"frozen_balance"`
		ExpireTime    int64 `json:"expire_time"`
	} `json:"frozen"`
	CreateTime            int64 `json:"create_time"`
	LatestOprationTime    int64 `json:"latest_opration_time"`
	LatestConsumeTime     int64 `json:"latest_consume_time"`
	LatestConsumeFreeTime int64 `json:"latest_consume_free_time"`
	AccountResource       struct {
		FrozenBalanceForEnergy struct {
			FrozenBalance int64 `json:"frozen_balance"`
			ExpireTime    int64 `json:"expire_time"`
		} `json:"frozen_balance_for_energy"`
		LatestConsumeTimeForEnergy int64 `json:"latest_consume_time_for_energy"`
	} `json:"account_resource"`
	AssetV2 []struct {
		Key   string `json:"key"`
		Value int64  `json:"value"`
	} `json:"assetV2"`
	FreeAssetNetUsageV2 []struct {
		Key   string `json:"key"`
		Value int64  `json:"value"`
	} `json:"free_asset_net_usageV2"`
}

// GetTrc20SymbolOut
type GetTrc20SymbolOut struct {
	Bytecode                   string `json:"bytecode"`
	ConsumeUserResourcePercent int    `json:"consume_user_resource_percent"`
	Name                       string `json:"name"`
	OriginAddress              string `json:"origin_address"`
	Abi                        struct {
		Entrys []struct {
			Outputs []struct {
				Type string `json:"type"`
			} `json:"outputs,omitempty"`
			Constant        bool   `json:"constant,omitempty"`
			Name            string `json:"name,omitempty"`
			StateMutability string `json:"stateMutability,omitempty"`
			Type            string `json:"type"`
			Inputs          []struct {
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"inputs,omitempty"`
		} `json:"entrys"`
	} `json:"abi"`
	OriginEnergyLimit int    `json:"origin_energy_limit"`
	ContractAddress   string `json:"contract_address"`
	CodeHash          string `json:"code_hash"`
}

// GetAssetIssueByID TRC10
type GetAssetIssueByID struct {
	OwnerAddress string `json:"owner_address"`
	Name         string `json:"name"`
	Abbr         string `json:"abbr"`
	TotalSupply  int64  `json:"total_supply"`
	TxrNum       int64  `json:"txr_num"`
	Precision    int64  `json:"precision"`
	Num          int64  `json:"num"`
	StartTime    int    `json:"start_time"`
	EndTime      int64  `json:"end_time"`
	Description  string `json:"description"`
	Url          string `json:"url"`
	Id           string `json:"id"`
}
