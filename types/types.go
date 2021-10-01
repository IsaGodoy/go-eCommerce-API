package types

import "encoding/json"

//Types for retrieve a product
type Variant struct {
	ID  int64  `json:"id"`
	SKU string `json:"sku"`
}

type Product struct {
	ID       int64     `json:"id"`
	Variants []Variant `json:"variants"`
}

type ProductRequest struct {
	Products []Product `json:"products"`
}

//Types for change a product
type VariantPut struct {
	ID               int64  `json:"id"`
	SKU              string `json:"sku"`
	Price            int64  `json:"price"`
	Compare_at_price int64  `json:"compare_at_price"`
}

type ProductPut struct {
	Variant VariantPut `json:"variant"`
}

//Other methods
func (pr *ProductRequest) ToJson() ([]byte, error) {
	return json.Marshal(pr)
}

func (pp *ProductPut) ToJson() ([]byte, error) {
	return json.Marshal(pp)
}
