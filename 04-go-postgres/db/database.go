package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const url = "postgres://postgres:2609@localhost:5432/postgres?sslmode=disable"

var db *sql.DB

func Connect() {
	conection, err := sql.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexion exitosa")
	db = conection
}

func Close() {
	defer db.Close()
}

func Ping() {
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Ping OK")
}

func CreateTable(schema string, name string) {
	if !ExistsTable(name) {
		_, err := db.Exec(schema)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func ExistsTable(tableName string) bool {
	sql := fmt.Sprintf("SELECT * FROM pg_catalog.pg_tables WHERE tablename = '%s'", tableName)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("Error", err)
	}
	return rows.Next()
}

// Polimorfismo a Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return rows, err
}
