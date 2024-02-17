package routes

import (
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
	err = c.repository.RealizarTransacao(&transaction)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao realizar transação: %v", err), http.StatusInternalServerError)
		return
	}
	transactionResponse := entity.TransactionResponse{
		Limite: 1000000,
		Saldo: 900000,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&transactionResponse)
	if err != nil {
		http.Error(w, "Erro ao retornar resultado da transação", http.StatusInternalServerError)
		return
	}
}