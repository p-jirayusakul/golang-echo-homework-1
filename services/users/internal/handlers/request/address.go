package request

type CreateAddressRequest struct {
	UserID   string `json:"userId" validate:"uuid4"`
	AddrType string `json:"addrType"`
	AddrNo   string `json:"addrNo"`
	Street   string `json:"street"`
	City     string `json:"city"`
	State    string `json:"state"`
}
