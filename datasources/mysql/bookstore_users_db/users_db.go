package bookstore_users_db

import (
	"database/sql"
)

var (
	Client *sql.DB
)

func init() {
	datasource := "root:root@tcp(127.0.01:3306)/bookstore_users"
	var err error // init error

	Client, err = sql.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}
	// verify connection
	if err = Client.Ping(); err != nil {
		panic(err)
	}
}
