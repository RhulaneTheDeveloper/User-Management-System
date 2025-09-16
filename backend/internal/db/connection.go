package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// ConnectMySQL creates a connection to MySQL
func ConnectMySQL() *sqlx.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASS"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB"),
	)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal("❌ MySQL connection error:", err)
	}
	log.Println("✅ Connected to MySQL")
	return db
}

// ConnectPostgres creates a connection to PostgreSQL
func ConnectPostgres() *sqlx.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASS"),
		os.Getenv("PG_DB"),
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal("❌ Postgres connection error:", err)
	}
	log.Println("✅ Connected to PostgreSQL")
	return db
}
