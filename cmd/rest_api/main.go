package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/zhenriquegomes/rinha-backend-2024/internal/infra/repository"
	"github.com/zhenriquegomes/rinha-backend-2024/internal/infra/rest/routes"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("Não foi possível se conectar do postgres: %v", err)
	}
	defer db.Close()
	clientRepository := repository.NewClientRepository(db)
	clientRouter := routes.NewClientRouter(clientRepository)
	mux := http.NewServeMux()
	mux.HandleFunc("POST /clientes/{id}/transacoes", clientRouter.RealizarTransacao)
	log.Println("Starting server on: http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}

