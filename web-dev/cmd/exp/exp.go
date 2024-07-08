package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DataBase string
	SSLMode  string
}

func (p *PostgresConfig) ConnectionString() string {
	return fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode = %s", p.Host, p.Port, p.User, p.Password, p.DataBase, p.SSLMode)
}

func main() {
	config := PostgresConfig{
		Host:     "localhost",
		Port:     5433,
		User:     "baloo",
		Password: "junglebook",
		DataBase: "lenslocked",
		SSLMode:  "disable",
	}

	db, err := sql.Open("pgx", config.ConnectionString())
	if err != nil {
		panic(err)
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database!")

	// create a table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT UNIQUE NOT NULL
		);

		CREATE TABLE IF NOT EXISTS orders (
			id SERIAL PRIMARY KEY,
			user_id INT REFERENCES users(id),
			amount INT,
			description TEXT
		);
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully created tables!")

	// insert some data...
	// name := "', ''); DROP TABLE users; --" // SQL Injection
	// name := "Bob2 Calhoun"
	// email := "bob2@email.com"
	// // query := fmt.Sprintf(`INSERT INTO users (name, email) VALUES ('%s', '%s');`, name, email)
	// // _, err = db.Exec(query)
	// row := db.QueryRow(`INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id;`, name, email)
	// var id int
	// err = row.Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Inserted a user!, id =", id)

	// querying record
	id := 1
	var name, email string
	row := db.QueryRow(`SELECT name, email FROM users WHERE id = $1;`, id)
	err = row.Scan(&name, &email)
	if err != nil {
		panic(err)
	}
	fmt.Println("Name:", name)

	userID := 1
	for i := 1; i <= 5; i++ {
		amount := i * 1000
		description := fmt.Sprintf("USB-C Adapter x %d", i)
		_, err = db.Exec(`INSERT INTO orders (user_id, amount, description) VALUES ($1, $2, $3);`, userID, amount, description)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Inserted orders!")

	// querying multiple records
	type Order struct {
		ID          int
		UserID      int
		Amount      int
		Description string
	}
	var orders []Order
	rows, err := db.Query(`SELECT id, user_id, amount, description FROM orders WHERE user_id = $1;`, userID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var order Order
		err = rows.Scan(&order.ID, &order.UserID, &order.Amount, &order.Description)
		if err != nil {
			panic(err)
		}
		orders = append(orders, order)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	fmt.Println("Orders:", orders)
}
