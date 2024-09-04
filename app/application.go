package app

import (
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartApplication() {
	mapUrls()
	router.Run("localhost:8080")
}
