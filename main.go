package main

import (
	"btpn-go/config"
	"btpn-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	config.ConnectDatabase()
	routes.SetupRoutes(r)

	// Menentukan port yang digunakan
	port := ":8080"
	r.Run(port)
}
