package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

func main() {
	// db_user:password@tcp(localhost:3306)/my_db
	db, err := sql.Open("mysql", "root:pass4sac%@tcp(localhost:3306)/test?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	{
		query := `CREATE TABLE users (
               id INT AUTO_INCREMENT,
               username TEXT NOT NULL,
               password TEXT NOT NULL,
               created_at DATETIME,
               PRIMARY KEY (id)
           );`
		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}
	{
		username := "johndoe"
		password := "secret"
		createdAt := time.Now()
		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt )
		if err != nil {
			log.Fatal(err)
		}
		id, err := result.LastInsertId()
		fmt.Println(id)
	}
	{
		fmt.Println("query a single user")
		var (
			id int
			username string
			password string
			createdAt time.Time
		)

		query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, username, password, createdAt)
	}

	{
		fmt.Println("query all users")
		type user struct {
			id        int
			username  string
			password  string
			createdAt time.Time
		}

		rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var users []user
		for rows.Next() {
			var u user

			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%#v", users)
	}

	{
		_, err := db.Exec(`delete from users where id = ?`, 1)
		if err != nil {
			log.Fatal(err)
		}
	}
}
