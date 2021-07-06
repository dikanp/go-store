package order

import (
	"AltaStore/models"
	// "AltaStore/models/product"
)

type Checkout struct {
	TotalAmount		int			`json:"total"`
	BankAccount		string	`json:"bank"`
	BankNumber		string	`json:"account_number"`
}

type CheckoutResponse struct {
	models.Response
	Data Checkout `json:"data"`
}
