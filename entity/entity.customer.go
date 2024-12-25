package entity

import (
	"time"
)

type Customer struct {
	CustomerId          string             `json:"customer_id"`
	CustomerCode        string             `json:"customer_code"`
	CustomerName        string             `json:"customer_name"`
	CustomerGender      string             `json:"customer_gender"`
	CustomerPhonenumber string             `json:"customer_phonenumber"`
	CustomerEmail       string             `json:"customer_email"`
	CustomerPassword    string             `json:"customer_password"`
	CustomerStatus      int                `json:"customer_status"`
	CustomerCreateAt    time.Time          `json:"customer_create_at"`
	CustomerUpdateAt    time.Time          `json:"customer_update_at"`
	Address             *[]CustomerAddress `json:"alamat"`
}

type CustomerResponse struct {
	CustomerId          string             `json:"customer_id"`
	CustomerCode        string             `json:"customer_code"`
	CustomerName        string             `json:"customer_name"`
	CustomerGender      string             `json:"customer_gender"`
	CustomerPhonenumber string             `json:"customer_phonenumber"`
	CustomerEmail       string             `json:"customer_email"`
	CustomerPassword    string             `json:"-"`
	CustomerStatus      int                `json:"customer_status"`
	CustomerCreateAt    time.Time          `json:"customer_create_at"`
	CustomerUpdateAt    time.Time          `json:"customer_update_at"`
	Address             *[]CustomerAddress `json:"alamat"`
}
