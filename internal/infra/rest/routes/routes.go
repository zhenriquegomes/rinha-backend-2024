package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/zhenriquegomes/rinha-backend-2024/internal/entity"
	"github.com/zhenriquegomes/rinha-backend-2024/internal/infra/repository"
)


type ClientRouter struct{
	repository *repository.ClientRepository
}

func NewClientRouter(repository *repository.ClientRepository) *ClientRouter {
	return &ClientRouter{repository: repository}
}

func (c *ClientRouter) RealizarTransacao(w http.ResponseWriter, r *http.Request) {
	strID := r.PathValue("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		http.Error(w, "Não foi possível ler o id do cliente", http.StatusUnprocessableEntity)
		return
	}
	var transaction entity.Transaction
	err = json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, "Não foi possível ler o body da requisição", http.StatusUnprocessableEntity)
		return
	}
	transaction.ClientID = id
	dataReferencia := time.Now()
	transaction.DataReferencia = &dataReferencia
	client, err := c.repository.ConsultarCliente(id)
	if err == sql.ErrNoRows {
		http.Error(w, "Cliente não encontrado", http.StatusNotFound)
		return
	} 
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao consultar o cliente: %v", err), http.StatusInternalServerError)
		return
	}
	err = client.ValidateTransaction(&transaction)
	if err == entity.ErrEstouroLimite {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	if err == entity.ErrTipoTransacaoInvalido {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	if err == entity.ErrDescricaoInvalida {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	if err != nil {
		http.Error(w, "Erro ao validar transação", http.StatusInternalServerError)
		return 
	}
	err = c.repository.RealizarTransacao(&transaction)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao realizar transação: %v", err), http.StatusInternalServerError)
		return
	}
	novoSaldo := client.SaldoAtual - transaction.Valor
	err = c.repository.AtualizarSaldo(novoSaldo, id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao atualizar saldo: %v", err), http.StatusInternalServerError)
		return
	}
	client, _ = c.repository.ConsultarCliente(id)
	transactionResponse := entity.TransactionResponse{
		Limite: client.Limite,
		Saldo: client.SaldoAtual,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&transactionResponse)
	if err != nil {
		http.Error(w, "Erro ao retornar resultado da transação", http.StatusInternalServerError)
		return
	}
}