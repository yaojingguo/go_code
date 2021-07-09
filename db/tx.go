package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", "sac:pass4sac@tcp(127.0.0.1:3306)/test")
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		panic(err.Error())
	}
	if _, err := tx.Exec(`update student set percent = 2 where name = 'xuyang'`); err != nil {
		_ = tx.Rollback()
		panic(err.Error())
	}
	fmt.Println("Sleeping...")
	time.Sleep(10 * time.Minute)
	if err := tx.Commit(); err != nil {
		panic(err.Error())
	}
}
