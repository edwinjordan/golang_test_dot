package entity

import "time"

type CustomerAddress struct {
	AddressId         string    `json:"address_id"`
	AddressCustomerId string    `json:"address_customer_id"`
	AddressText       string    `json:"address_text"`
	AddressName       string    `json:"address_name"`
	AddressPostalCode string    `json:"address_postal_code"`
	AddressCreateAt   time.Time `json:"address_create_at"`
	AddressUpdateAt   time.Time `json:"address_update_at"`
}

type CustomerAddressResponse struct {
	AddressId         string    `json:"address_id"`
	AddressCustomerId string    `json:"address_customer_id"`
	AddressText       string    `json:"address_text"`
	AddressName       string    `json:"address_name"`
	AddressPostalCode string    `json:"address_postal_code"`
	AddressCreateAt   time.Time `json:"address_create_at"`
	AddressUpdateAt   time.Time `json:"address_update_at"`
}
