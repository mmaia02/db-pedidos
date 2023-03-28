package main

import (
    "database/sql"
    "fmt"
    "log"

    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

func main() {
    db, err := sqlx.Connect("postgres", "user=postgres password=password dbname=pedidosdb sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
}

    defer db.Close()

    type Pedido struct {
        ID           int64   `db:"id"`
        NomeCliente  string  `db:"nome_cliente"`
        Valor        float64 `db:"valor"`
    }

    func createPedido(db *sqlx.DB, p *Pedido) error {
        query := `INSERT INTO pedidos (nome_cliente, valor) VALUES ($1, $2) RETURNING id`
        return db.QueryRowx(query, p.NomeCliente, p.Valor).Scan(&p.ID)
    }
    
    func getPedido(db *sqlx.DB, id int64) (*Pedido, error) {
        p := &Pedido{}
        err := db.Get(p, "SELECT * FROM pedidos WHERE id=$1", id)
    }