package entity

import "time"

type CustomerOrder struct {
	OrderId         string    `json:"order_id"`
	OrderCustomerId string    `json:"order_customer_id"`
	OrderInvNumber  string    `json:"order_inv_number"`
	OrderAddressId  string    `json:"order_address_id"`
	OrderTotalItem  int       `json:"order_total_item"`
	OrderSubtotal   float64   `json:"order_subtotal"`
	OrderDiscount   float64   `json:"order_discount"`
	OrderTotal      float64   `json:"order_total"`
	OrderNotes      string    `json:"order_notes"`
	OrderStatus     int       `json:"order_status"`
	OrderCreateAt   time.Time `json:"order_create_at"`
	OrderDetail     *[]CustomerOrderDetail
	Customer        *CustomerResponse `json:"customer"`
	//	Address         *CustomerAddress  `json:"address"`
}
