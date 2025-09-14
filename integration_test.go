package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
	"testing"
)

func setupTestDB(t *testing.T) *sqlx.DB {
	dsn := os.Getenv("TEST_DATABASE_DSN")
	if dsn == "" {
		t.Skip("TEST_DATABASE_DSN not set; skip integration test")
	}
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		t.Fatalf("failed connect to db: %v", err)
	}
	// очистка таблиц, пример:
	// db.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE")
	return db
}

func TestDBExample(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	// пример: проверим, что можем выполнить простой запрос
	var now string
	if err := db.Get(&now, "SELECT now()::text"); err != nil {
		t.Fatalf("db query failed: %v", err)
	}
	t.Logf("DB responded: %s", now)
}
