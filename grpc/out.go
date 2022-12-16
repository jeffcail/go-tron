package grpc

type GetTrc10TokenListOut struct {
	ID           string `json:"id"`
	OwnerAddress string `json:"owner_address"`
	Name         string `json:"name"`
	Abbr         string `json:"abbr"`
	Decimal      int32  `json:"decimal"`
}
