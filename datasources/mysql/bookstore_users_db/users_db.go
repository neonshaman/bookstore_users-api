package bookstore_users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	bookstore_users_username = "bookstore_users_username"
	bookstore_users_password = "bookstore_users_password"
	bookstore_users_host     = "bookstore_users_host"
	bookstore_users_schema   = "bookstore_users_schema"
)

var (
	Client *sql.DB
	// grab from env for now, secrets managers of some kind later
	username = os.Getenv(bookstore_users_username)
	password = os.Getenv(bookstore_users_password)
	host     = os.Getenv(bookstore_users_host)
	schema   = os.Getenv(bookstore_users_schema)
)

func init() {
	datasource := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, schema)
	var err error // init error

	Client, err = sql.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}
	// verify connection
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Successfully connected to bookstore_users")
}
