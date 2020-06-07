package helpers

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "@123elvischuks"
	dbname   = "ecommerce"
)

func InitDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s"+" password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db
}

func InitTables() error {
	db := InitDB()
	defer db.Close()

	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS products(id SERIAL PRIMARY KEY, category VARCHAR, name VARCHAR, quantity VARCHAR, price VARCHAR,image VARCHAR)")

	_, err := db.Exec(query)
	query = fmt.Sprintf("CREATE TABLE IF NOT EXISTS categories(id SERIAL PRIMARY KEY, name VARCHAR)")

	_, err = db.Exec(query)

	if err != nil {
		return err
	}
	return nil
}

func DropTables() error {
	db := InitDB()
	defer db.Close()

	query := fmt.Sprintf("DROP TABLE IF EXISTS products")

	_, err := db.Exec(query)

	if err != nil {
		return err
	}
	return nil
}
