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