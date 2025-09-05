package pkg

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	DBAddr string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

func NewDatabase(addr, port, user, pass, name string) *Database {
	return &Database{
		DBAddr: addr,
		DBPort: port,
		DBUser: user,
		DBPass: pass,
		DBName: name,
	}
}

func (db *Database) Connect() *pgxpool.Pool {
	address := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		db.DBUser, db.DBPass, db.DBAddr, db.DBPort, db.DBName)

	pool, err := pgxpool.New(context.Background(), address)
	if err != nil {
		println(fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err))
		os.Exit(1)
	}

	_, err = pool.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS movies (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			description TEXT
		)
	`)
	if err != nil {
		println(fmt.Fprintf(os.Stderr, "Unable to create table: %v\n", err))
		os.Exit(1)
	}

	return pool
}
