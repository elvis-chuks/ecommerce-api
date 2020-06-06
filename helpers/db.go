package helpers

import (
	"fmt"
	"database/sql"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "@123elvischuks"
	dbname = "ecommerce"
)

func InitDB() *sql.DB{
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s"+" password=%s dbname=%s sslmode=disable",
	host,port,user,password,dbname)

	db, err := sql.Open("postgres",psqlInfo)
	if err != nil{
		panic(err)
	}

	return db
}

func InitTables() error {
	db := InitDB()
	defer db.Close()

	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS products(id SERIAL PRIMARY KEY, category VARCHAR, name VARCHAR, quantity VARCHAR, price VARCHAR)")

	_, err := db.Exec(query)

	if err != nil {
		return err
	}
	return nil
}