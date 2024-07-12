package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type user_info struct {
	ID          int64
	tg_nickname string
	balance     float32
}

func main() {

	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}
	fmt.Println("DBUSER")
	fmt.Println("DBPASS")

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("DBUSER")
	fmt.Println("DBPASS")

	UserID, err := addUser(user_info{
		tg_nickname: "andrewthefuckingambler",
		balance:     1000,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID of added user: %v\n", UserID)

}

func addUser(NewUser user_info) (int64, error) {
	result, err := db.Exec("INSERT INTO user_info (tg_nickname, balance) VALUES (?, ?)", NewUser.tg_nickname, NewUser.balance)
	if err != nil {
		return 0, fmt.Errorf("addUser: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addUser: %v", err)
	}
	return id, nil
}
