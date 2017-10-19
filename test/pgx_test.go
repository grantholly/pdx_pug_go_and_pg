package main

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/jackc/pgx/stdlib"
)

const (
	DB_HOST = "192.168.0.100"
	DB_PORT = "5432"
	DB_USER = "postgres"
	DB_PASS = "postgres"
	DB_NAME = "demo"
)

var dbinfo = fmt.Sprintf("host=%s port=%s "+
	"user=%s password=%s "+
	"dbname=%s sslmode=disable",
	DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)

//benchmark helper
func db_connect() *sql.DB {
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("couldn't connect to DB")
	}
	return db
}

// no prepared statement
func benchmarkPgxInsert(b *testing.B) {
	db := db_connect()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		db.Exec("INSERT INTO public.tickets VALUES ($1, $2, $3)",
			n, "action", "closed")
	}
}

func benchmarkPgxUpdate(b *testing.B) {

}

func benchmarkPgxDelete(b *testing.B) {

}

func benchmarkPgxSelect(b *testing.B) {

}

//benchmarks
func BenchmarkPgxInsert1000 (b *testing.B) { benchmarkPgxInsert(b) }
