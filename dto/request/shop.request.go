package request

type ShopRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type TypeProductRequest struct {
	ShopID uint   `json:"shopId"`
	Hastag string `json:"hastag"`
	Name   string `json:"name"`
}
