package entity

type Product struct {
	ID              int             `json:"id"`
	Code            string          `json:"code"`
	Name            string          `json:"name"`
	Description     string          `json:"description"`
	Price           float32         `json:"price"`
	ProductCategory ProductCategory `json:"product_category"`
}
