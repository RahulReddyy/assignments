package models

type Customer struct {
	CName  string `json:"name"`
	CAccNo int    `json:"account_Number"`
	CSSN   int    `json:"ssn"`
	Amount int    `json:"amount"`
}

type Bank struct {
	GenCus []Customer
}
