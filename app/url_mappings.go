package app

import "github.com/neonshaman/bookstore_users-api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
