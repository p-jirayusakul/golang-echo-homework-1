package entities

import "github.com/google/uuid"

type Address struct {
	AddressId uuid.UUID `json:"addressId"`
	UserID    uuid.UUID `json:"userId"`
	AddrType  string    `json:"addrType"`
	AddrNo    string    `json:"addrNo"`
	Street    string    `json:"street"`
	City      string    `json:"city"`
	State     string    `json:"state"`
}
