package main

import (
	"fmt"
	"loloMart/models"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)

}

func main() {
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected！")

	// Create a new user
	us := models.UserService{
		DB: db,
	}

	user, err := us.Create("lawrence4@li.com", "lawrence123")
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

	// Create a table
	// _, err = db.Exec(`
	// CREATE TABLE IF NOT EXISTS users (
	// 	id SERIAL PRIMARY KEY,
	// 	name TEXT,
	// 	email TEXT UNIQUE NOT NULL
	// );

	// CREATE TABLE IF NOT EXISTS orders (
	// 	id SERIAL PRIMARY KEY,
	// 	user_id INT NOT NULL,
	// 	amount INT,
	// 	description TEXT
	// );
	// `)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Table created！")

	//insert some data
	// name := "Larry Bird"
	// email := "larry@gmail.com"

	// row := db.QueryRow(`
	// INSERT INTO users (name, email)
	// VALUES ($1,$2) RETURNING id;`, name, email)

	// var id int
	// err = row.Scan(&id)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("User created！id =", id)

	// id := 4
	// row := db.QueryRow(`
	// SELECT name, email
	// FROM users
	// WHERE id = $1;`, id)

	// var name, email string

	// err = row.Scan(&name, &email)

	// if err == sql.ErrNoRows {
	// 	fmt.Println("No row found！")
	// }

	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("User information: name = %s, email = %s\n", name, email)

	// for i := 0; i <= 5; i++ {
	// 	amount := i * 100
	// 	desc := fmt.Sprintf("Fake Order %d", i)
	// 	_, err = db.Exec(`
	// 	INSERT INTO orders (user_id, amount, description)
	// 	VALUES ($1,$2,$3);`, userID, amount, desc)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// fmt.Println("Created fake orders.")

	// type Order struct {
	// 	Id          int
	// 	UserID      int
	// 	Amount      int
	// 	Description string
	// }

	// var orders []Order

	// userID := 1

	// rows, err := db.Query(`
	// SELECT id, amount, description
	// FROM orders
	// WHERE user_id = $1;`, userID)

	// if err != nil {
	// 	panic(err)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var order Order
	// 	order.UserID = userID
	// 	err = rows.Scan(&order.Id, &order.Amount, &order.Description)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	orders = append(orders, order)
	// }

	// err = rows.Err()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Orders: ", orders)
}
