package app

import (
	"github.com/neonshaman/bookstore_users-api/controllers/ping"
	"github.com/neonshaman/bookstore_users-api/controllers/users"
)

func mapUrls() {
	// Health check
	router.GET("/ping", ping.Ping)
	// Users
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/find", users.FindUser)

}
