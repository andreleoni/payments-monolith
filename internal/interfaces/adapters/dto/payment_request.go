package dto

// Address represents the address of the user.
type Address struct {
	Street     string `json:"street"`
	Zipcode    string `json:"zipcode"`
	Number     string `json:"number"`
	Complement string `json:"complement"`
}

// CreditCard represents the credit card details.
type CreditCard struct {
	Number  string `json:"Number"`
	CVV     string `json:"cvv"`
	Expires string `json:"expires"`
}

// User represents the user details including their address.
type User struct {
	FullName string  `json:"full_name"`
	CPF      string  `json:"cpf"`
	Address  Address `json:"address"`
}

// PaymentRequest represents the structure of the request body.
type PaymentRequest struct {
	Identifier string     `json:"identifier"`
	CreditCard CreditCard `json:"credit_card"`
	User       User       `json:"user"`
	ValueCents uint64     `json:"value_cents"`
}
