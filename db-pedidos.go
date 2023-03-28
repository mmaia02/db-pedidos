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