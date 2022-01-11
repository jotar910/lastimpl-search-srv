package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pass123" // temporary
	dbname   = "temporary"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection success")
	rows, err := db.Query(`
	SELECT users.id, users.name, movements.description
	FROM users
	INNER JOIN movements ON users.id = movements.user_id
	`)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name, description string
		err := rows.Scan(&id, &name, &description)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("id: ", id, "name: ", name, "description: ", description)
	}
	if err := rows.Err(); err != nil {
		log.Fatalln(err)
	}
	log.Println("Terminated with success")
}
