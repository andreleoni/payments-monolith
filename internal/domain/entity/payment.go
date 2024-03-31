package entity

type Payment struct {
	ID                        string `json:"id"`
	Identifier                string `json:"identifier"`
	UserFullName              string `json:"full_name"`
	UserCPF                   string `json:"cpf"`
	CreditCardNumber          string `json:"Number"`
	CreditCardCVV             string `json:"cvv"`
	CreditCardExpires         string `json:"expires"`
	AddressStreet             string `json:"street"`
	AddressZipcode            string `json:"zipcode"`
	AddressNumber             string `json:"number"`
	AddressComplement         string `json:"complement"`
	CreatedAt                 string `json:"created_at"`
	ValueCents                uint64 `json:"value_cents"`
	ExternalServiceIdentifier string `json:"external_service_identifier"`
	State                     string `json:"state"`
	Error                     string `json:"error"`
}
