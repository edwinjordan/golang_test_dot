package entity

type CustomerOrderDetail struct {
	OrderDetailId       string  `json:"order_detail_id"`
	OrderDetailParentId string  `json:"order_detail_parent_id"`
	OrderDetailProduct  string  `json:"order_detail_product"`
	OrderDetailQty      int     `json:"order_detail_qty"`
	OrderDetailPrice    float64 `json:"order_detail_price"`
	OrderDetailSubtotal float64 `json:"order_detail_subtotal"`
}

type ViewOrderDetail struct {
	OrderDetailId       string  `json:"order_detail_id"`
	OrderDetailParentId string  `json:"order_detail_parent_id"`
	OrderDetailProduct  string  `json:"order_detail_product"`
	OrderDetailQty      int     `json:"order_detail_qty"`
	OrderDetailPrice    float64 `json:"order_detail_price"`
	OrderDetailSubtotal float64 `json:"order_detail_subtotal"`
	//ProductName         string  `json:"product_name"`
}
