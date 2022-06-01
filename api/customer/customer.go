package customer

import "time"

type CustomerData struct {
	CustomerID           string              `json:"customer_id"`
	Name                 string              `json:"name"`
	Surname              string              `json:"surname"`
	Mail                 string              `json:"mail"`
	Phone                string              `json:"phone"`
	Age                  time.Time           `json:"age"`
	IdentificationNumber string              `json:"identification_number"` //TODO Should masked
	Status               string              `json:"status"`
	CreateDate           time.Time           `json:"create_date"`
	Address              CustomerAddressData `json:"address"`
}

type CustomerAddressData struct {
	Country     string `json:"country"`
	City        string `json:"city"`
	District    string `json:"district"`
	PostalCode  string `json:"postal_code"`
	LongAddress string `json:"long_address"`
}
