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

type Cliente struct {
	ID int `json:"id"`
	Limite int `json:"limite"`
	SaldoInicial int `json:"saldo_inicial"`
	SaldoAtual int `json:"saldo_atual"`
}