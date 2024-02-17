package repository

import (
	"database/sql"

	"github.com/zhenriquegomes/rinha-backend-2024/internal/entity"
)

type ClientRepository struct {
	db *sql.DB
}

func NewClientRepository(db *sql.DB) *ClientRepository {
	return &ClientRepository{db: db}
}

func (c *ClientRepository) RealizarTransacao(transaction *entity.Transaction) error {
	stmt, err := c.db.Prepare(`
		INSERT INTO public.transacoes
			(client_id, valor, tipo, descricao, dt_referencia)
		VALUES
			($1, $2, $3, $4, $5);
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		transaction.ClientID,
		transaction.Valor,
		transaction.Tipo,
		transaction.Descricao,
		transaction.DataReferencia,
	)
	if err != nil {
		return err
	}
	return nil
}