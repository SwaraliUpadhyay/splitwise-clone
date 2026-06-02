package server

import (
	"splitwise-clone/routes"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	routes.Register(router)
	router.Run(":8080")
}
