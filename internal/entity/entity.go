package entity

import (
	"errors"
	"time"
)


var (
	ErrEstouroLimite = errors.New("a transação estourou o limite do cliente")
	ErrTipoTransacaoInvalido = errors.New("tipo de transação inválido")
	ErrDescricaoInvalida = errors.New("descrição de transação inválida")
)

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

func (c *Cliente) abs(n int) int {
	if n > 0 {
		return n 
	} else {
		return -n
	}
}

func (c *Cliente) ValidateTransaction(transaction *Transaction) error {
	if transaction.Tipo == "d" && c.abs(c.SaldoAtual - transaction.Valor) > c.Limite {
		return ErrEstouroLimite
	}
	if transaction.Tipo != "c" && transaction.Tipo != "d" {
		return ErrTipoTransacaoInvalido
	}
	if transaction.Descricao == "" || len(transaction.Descricao) > 10 {
		return ErrDescricaoInvalida
	}
	return nil
}
