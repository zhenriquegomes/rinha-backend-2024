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

func (c *ClientRepository) ConsultarCliente(id int) (*entity.Cliente, error) {
	stmt, err := c.db.Prepare(`
		SELECT id, limite, saldo_inicial, saldo_atual FROM public.clientes WHERE id = $1
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var cliente entity.Cliente
	err = stmt.QueryRow(id).Scan(&cliente.ID, &cliente.Limite, &cliente.SaldoInicial, &cliente.SaldoAtual)
	if err != nil {
		return nil, err
	}	
	return &cliente, nil
}

func (c *ClientRepository) AtualizarSaldo(novo_saldo, cliente_id int) error {
	stmt, err := c.db.Prepare(`
		UPDATE public.clientes
		SET saldo_atual = $1
		WHERE id = $2
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(novo_saldo, cliente_id)
	if err != nil {
		return err
	}
	return nil
}