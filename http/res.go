package http

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
