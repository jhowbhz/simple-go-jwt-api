package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDatabase() {

	var err error

	db, err = sql.Open("sqlite3", "./database/banco.db")

	if err != nil {
		log.Fatalf("Erro ao abrir o banco de dados: %v", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL
		)
	`)

	if err != nil {
		log.Fatalf("Erro ao criar a tabela users: %v", err)
	}

	log.Println("Banco de dados inicializado com sucesso!")
}
