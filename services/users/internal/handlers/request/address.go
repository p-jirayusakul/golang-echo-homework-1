package request

type CreateAddressRequest struct {
	UserID   string `json:"userId" validate:"uuid4"`
	AddrType string `json:"addrType"`
	AddrNo   string `json:"addrNo"`
	Street   string `json:"street"`
	City     string `json:"city"`
	State    string `json:"state"`
}

type FindAddressRequest struct {
	AddressId string `param:"address_id" validate:"uuid4"`
}

type UpdateAddressRequest struct {
	AddressId string `json:"addressId" validate:"uuid4"`
	AddrType  string `json:"addrType"`
	AddrNo    string `json:"addrNo"`
	Street    string `json:"street"`
	City      string `json:"city"`
	State     string `json:"state"`
}

type DeleteAddressRequest struct {
	AddressId string `param:"address_id" validate:"uuid4"`
}
