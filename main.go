package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	//gets an engine (http server) with logger and recovery
	server := gin.Default()

	//pass pointer
	routes.RegisterRoutes(server)
	server.Run(":8080") // start the server to listen to requests on local host

}
