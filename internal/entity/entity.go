package entity

import "time"

type Transaction struct {
	ClientID int
	Valor int `json:"valor"`
	Tipo string `json:"tipo"`
	Descricao string `json:"descricao"`
	DataReferencia *time.Time
}

type TransactionResponse struct {
	Limite int `json:"limite"`
	Saldo int `json:"saldo"`
}